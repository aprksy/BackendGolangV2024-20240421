package handler_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	repo "github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func TestPostEstate(t *testing.T) {
	testCases := []struct {
		name    string
		jsonStr string
	}{
		{name: "invalid", jsonStr: `{"width":, "length":}`},
		{name: "invalid", jsonStr: `{"width":-1, "length":1}`},
		{name: "invalid", jsonStr: `{"width":1, "length":-1}`},
		{name: "invalid", jsonStr: `{"width":50001, "length":1}`},
		{name: "invalid", jsonStr: `{"width":1, "length":50001}`},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	mockRepo := repo.NewMockRepositoryInterface(ctrl)

	server := &handler.Server{
		Repository: mockRepo,
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(tc.jsonStr))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			err := server.PostEstate(ctx)
			assert.NotNil(t, err)
		})
	}

	userJSON := `{"length": 10, "width": 20}`
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().CreateEstate(ctx.Request().Context(), repo.CreateEstateInput{Length: 10, Width: 20}).
		Return(repo.CreateEstateOutput{Id: uuid.New()}, nil).Times(1)

	err := server.PostEstate(ctx)
	assert.Nil(t, err)
}

func TestPostEstateIdTree(t *testing.T) {
	testCases := []struct {
		name    string
		jsonStr string
	}{
		{name: "invalid", jsonStr: `{"x":, "y":, "height":}`},
		{name: "invalid", jsonStr: `{"x":-1, "y":1, "height":10}`},
		{name: "invalid", jsonStr: `{"x":1, "y":-1, "height":10}`},
		{name: "invalid", jsonStr: `{"x":50001, "y":1, "height":10}`},
		{name: "invalid", jsonStr: `{"x":1, "y":50001, "height":10}`},
		{name: "invalid", jsonStr: `{"x":1, "y":1, "height":-1}`},
		{name: "invalid", jsonStr: `{"x":1, "y":1, "height":31}`},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	mockRepo := repo.NewMockRepositoryInterface(ctrl)

	server := &handler.Server{
		Repository: mockRepo,
	}

	id := uuid.New()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/tree", id.String()), strings.NewReader(tc.jsonStr))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			err := server.PostEstateIdTree(ctx, id)
			assert.NotNil(t, err)
		})
	}

	userJSON := `{"x": 10, "y": 20, "height":20}`
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/tree", id.String()), strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	mockRepo.EXPECT().SetTreeHeight(ctx.Request().Context(), repo.SetTreeHeightInput{Id: id, X: 10, Y: 20, Height: 20}).
		Return(repo.SetTreeHeightOutput{Id: id}, nil).Times(1)

	err := server.PostEstateIdTree(ctx, id)
	assert.Nil(t, err)
}

func TestGetEstateIdStats(t *testing.T) {
	id_ok := uuid.New()
	id_nonexist := uuid.New()
	testCases := []struct {
		name          string
		id            uuid.UUID
		expectErr     error
		expectHttpErr *echo.HTTPError
	}{
		{name: "invalid", id: id_nonexist, expectErr: fmt.Errorf("NOTFOUND"), expectHttpErr: echo.ErrNotFound},
		{name: "invalid", id: id_nonexist, expectErr: fmt.Errorf("INTERNAL SERVER ERROR"), expectHttpErr: echo.ErrInternalServerError},
		{name: "invalid", id: id_ok, expectErr: nil, expectHttpErr: nil},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	mockRepo := repo.NewMockRepositoryInterface(ctrl)

	server := &handler.Server{
		Repository: mockRepo,
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s/stats", tc.id.String()), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			mockRepo.EXPECT().GetEstateStats(ctx.Request().Context(), repo.GetEstateStatsInput{Id: tc.id}).
				Return(repo.GetEstateStatsOutput{Count: 0, Min: 0, Max: 0, Median: 0}, tc.expectErr).Times(1)

			err := server.GetEstateIdStats(ctx, tc.id)
			if tc.expectHttpErr == nil {
				assert.Equal(t, nil, err)
			} else {
				assert.Equal(t, tc.expectHttpErr, err)
			}
		})
	}
}

func TestGetEstateIdDronePlan(t *testing.T) {
	id_ok := uuid.New()
	id_nonexist := uuid.New()
	testCases := []struct {
		name          string
		id            uuid.UUID
		maxDistance   int
		expectErr     error
		expectHttpErr *echo.HTTPError
	}{
		{name: "invalid", id: id_nonexist, maxDistance: 50, expectErr: fmt.Errorf("NOTFOUND"), expectHttpErr: echo.ErrNotFound},
		{name: "invalid", id: id_nonexist, maxDistance: 50, expectErr: fmt.Errorf("INTERNAL SERVER ERROR"), expectHttpErr: echo.ErrInternalServerError},
		{name: "invalid", id: id_ok, maxDistance: 50, expectErr: nil, expectHttpErr: nil},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	mockRepo := repo.NewMockRepositoryInterface(ctrl)

	server := &handler.Server{
		Repository: mockRepo,
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s/drone-plan?max-distance=%d", tc.id.String(), tc.maxDistance), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			mockRepo.EXPECT().PlanDronePatrol(ctx.Request().Context(), repo.PlanDronePatrolInput{Id: tc.id, MaxDistance: tc.maxDistance}).
				Return(repo.PlanDronePatrolOutput{Distance: 10, Rest: repo.Rest{X: 1, Y: 1}}, tc.expectErr).Times(1)

			err := server.GetEstateIdDronePlan(ctx, tc.id, generated.GetEstateIdDronePlanParams{MaxDistance: &tc.maxDistance})
			if tc.expectHttpErr == nil {
				assert.Equal(t, nil, err)
			} else {
				assert.Equal(t, tc.expectHttpErr, err)
			}
		})
	}
}
