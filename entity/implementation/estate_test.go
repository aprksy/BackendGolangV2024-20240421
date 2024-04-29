package entity_impl_test

import (
	"testing"

	entity_impl "github.com/SawitProRecruitment/UserService/entity/implementation"
	"github.com/stretchr/testify/assert"
)

func TestNewEstateFail(t *testing.T) {
	testCases := []struct {
		name              string
		width             int
		length            int
		expenctedInstance *entity_impl.Estate
		expectedWidth     int
		expectedLength    int
		expectedError     error
	}{
		// failure cases
		{name: "width is less than 1 ", width: -1, length: 10, expenctedInstance: nil, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidEstateDimension)},
		{name: "width is more than 50000", width: 67100, length: 10, expenctedInstance: nil, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidEstateDimension)},
		{name: "length is less than 1 ", width: 10, length: -3, expenctedInstance: nil, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidEstateDimension)},
		{name: "length is more than 50000", width: 10, length: 67100, expenctedInstance: nil, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidEstateDimension)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, err := entity_impl.NewEstate(tc.width, tc.length)
			if assert.Error(t, err) {
				assert.Equal(t, tc.expectedError, err, "error should be ErrInvalidEstateDimension")
				assert.Nil(t, estate, "estate instance should be nil")
			}
		})
	}
}

func TestNewEstateSuccess(t *testing.T) {
	testCases := []struct {
		name           string
		width          int
		length         int
		expectedWidth  int
		expectedLength int
	}{
		// success cases
		{name: "width is in range (1-50000)", width: 13, length: 10, expectedWidth: 13, expectedLength: 10},
		{name: "length is in range (1-50000)", width: 10, length: 14, expectedWidth: 10, expectedLength: 14},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, err := entity_impl.NewEstate(tc.width, tc.length)

			assert.NotNil(t, estate, "estate should not be nil")
			assert.Nil(t, err, "error should be nil")
			assert.Equal(t, tc.expectedWidth, estate.Width(), "width is incorrect")
			assert.Equal(t, tc.expectedLength, estate.Length(), "length is incorrect")
			assert.NotEmpty(t, estate.Id(), "id should not be empty")
		})
	}
}

func TestGetPlotFail(t *testing.T) {
	const (
		width  = 10 // y
		length = 20 // x
	)

	testCases := []struct {
		name          string
		x             int
		y             int
		expectedError error
	}{
		{name: "x is below minimum", x: 0, y: 1, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidCoordinates)},
		{name: "y is below minimum", x: 1, y: 0, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidCoordinates)},
		{name: "x is over maximum", x: 30, y: 1, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidCoordinates)},
		{name: "y is over maximum", x: 1, y: 21, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidCoordinates)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, _ := entity_impl.NewEstate(width, length)
			plot, err := estate.GetPlot(tc.x, tc.y)
			if assert.Error(t, err) {
				assert.Equal(t, tc.expectedError, err, "err msg is incorrect")
				assert.Nil(t, plot, "plot should be nil")
			}
		})
	}
}

func TestGetPlotSuccess(t *testing.T) {
	const (
		width  = 10 // y
		length = 20 // x
	)

	testCases := []struct {
		name       string
		x          int
		y          int
		expectedId string
	}{
		// edge cases
		{name: "x = 1, y = 1", x: 1, y: 1, expectedId: "1:1"},
		{name: "x = 20, y = 10", x: 20, y: 10, expectedId: "20:10"},
		{name: "x = 1, y = 10", x: 1, y: 10, expectedId: "1:10"},
		{name: "x = 20, y = 1", x: 20, y: 1, expectedId: "20:1"},
		// inclusive cases
		{name: "x = 14, y = 3", x: 14, y: 3, expectedId: "14:3"},
		{name: "x = 3, y = 6", x: 3, y: 6, expectedId: "3:6"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, _ := entity_impl.NewEstate(width, length)
			plot, _ := estate.GetPlot(tc.x, tc.y)
			assert.Equal(t, tc.expectedId, plot.Id(), "id is incorrect")
		})
	}
}

func TestGetSetTreeHeightFromEstateFail(t *testing.T) {
	const (
		width  = 10 // y
		length = 20 // x
	)

	testCases := []struct {
		name   string
		x      int
		y      int
		height int
	}{
		{name: "x is below minimum", x: 0, y: 1, height: 10},
		{name: "y is below minimum", x: 1, y: 0, height: 10},
		{name: "x is over maximum", x: 30, y: 1, height: 10},
		{name: "y is over maximum", x: 1, y: 21, height: 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, _ := entity_impl.NewEstate(width, length)
			err := estate.SetTreeHeight(tc.x, tc.y, tc.height)
			if assert.Error(t, err) {
				assert.Equal(t, entity_impl.NewErr(entity_impl.ErrInvalidCoordinates), err, "err msg is incorrect")
				height, err1 := estate.GetTreeHeight(tc.x, tc.y)
				if assert.Error(t, err1) {
					assert.Equal(t, entity_impl.NewErr(entity_impl.ErrInvalidCoordinates), err1)
					assert.Equal(t, -1, height, "height should be -1")
				}
			}
		})
	}

	testCases = []struct {
		name   string
		x      int
		y      int
		height int
	}{
		{name: "height is below minimum", x: 1, y: 1, height: -1},
		{name: "height is below maximum", x: 1, y: 1, height: 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, _ := entity_impl.NewEstate(width, length)
			err := estate.SetTreeHeight(tc.x, tc.y, tc.height)
			if assert.Error(t, err) {
				assert.Equal(t, entity_impl.NewErr(entity_impl.ErrInvalidTreeHeight), err, "err msg is incorrect")
				height, err1 := estate.GetTreeHeight(tc.x, tc.y)
				assert.Nil(t, err1, "err1 should be nil")
				assert.Equal(t, 0, height, "height should be 0")
			}
		})
	}
}

func TestGetSetTreeHeightFromEstateSuccess(t *testing.T) {
	const (
		width  = 10 // y
		length = 20 // x
	)

	testCases := []struct {
		name   string
		x      int
		y      int
		height int
	}{
		// edge cases
		{name: "x = 1, y = 1", x: 1, y: 1, height: 1},
		{name: "x = 20, y = 10", x: 20, y: 10, height: 30},
		{name: "x = 1, y = 10", x: 1, y: 10, height: 1},
		{name: "x = 20, y = 1", x: 20, y: 1, height: 30},
		// inclusive cases
		{name: "x = 14, y = 3", x: 14, y: 3, height: 15},
		{name: "x = 3, y = 6", x: 3, y: 6, height: 15},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, _ := entity_impl.NewEstate(width, length)
			err := estate.SetTreeHeight(tc.x, tc.y, tc.height)
			assert.Nil(t, err, "err should be nil")
			height, err1 := estate.GetTreeHeight(tc.x, tc.y)
			assert.Nil(t, err1, "err should be nil")
			assert.Equal(t, tc.height, height, "height should be the same")
		})
	}
}
