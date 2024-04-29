package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// (POST /estate)
func (s *Server) PostEstate(ctx echo.Context) error {
	var input repository.CreateEstateInput
	err := json.NewDecoder(ctx.Request().Body).Decode(&input)
	if err != nil {
		return echo.ErrBadRequest
	}

	if input.Length <= 0 || input.Length > 50000 {
		return echo.ErrBadRequest
	}

	if input.Width <= 0 || input.Width > 50000 {
		return echo.ErrBadRequest
	}

	resp, err := s.Repository.CreateEstate(ctx.Request().Context(), input)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, resp)
}

// (GET /estate/{id}/drone-plan)
func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id openapi_types.UUID, params generated.GetEstateIdDronePlanParams) error {
	md := -1
	if params.MaxDistance != nil {
		md = *params.MaxDistance
	}
	input := repository.PlanDronePatrolInput{
		Id:          id,
		MaxDistance: md,
	}

	resp, err := s.Repository.PlanDronePatrol(ctx.Request().Context(), input)
	if err != nil {
		if err.Error() == "NOTFOUND" {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, resp)
}

// (GET /estate/{id}/stats)
func (s *Server) GetEstateIdStats(ctx echo.Context, id openapi_types.UUID) error {
	input := repository.GetEstateStatsInput{
		Id: id,
	}
	resp, err := s.Repository.GetEstateStats(ctx.Request().Context(), input)
	if err != nil {
		if err.Error() == "NOTFOUND" {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, resp)
}

// (POST /estate/{id}/tree)
func (s *Server) PostEstateIdTree(ctx echo.Context, id openapi_types.UUID) error {
	var input repository.SetTreeHeightInput
	err := json.NewDecoder(ctx.Request().Body).Decode(&input)
	if err != nil {
		return echo.ErrBadRequest
	}

	if input.X <= 0 || input.X > 50000 {
		return echo.ErrBadRequest
	}

	if input.Y <= 0 || input.Y > 50000 {
		return echo.ErrBadRequest
	}

	if input.Height <= 0 || input.Height > 30 {
		return echo.ErrBadRequest
	}

	input.Id = id
	resp, err := s.Repository.SetTreeHeight(ctx.Request().Context(), input)
	if err != nil {
		if err.Error() == "NOTFOUND" {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, resp)
}
