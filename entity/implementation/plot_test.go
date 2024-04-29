package entity_impl_test

import (
	"testing"

	entity_impl "github.com/SawitProRecruitment/UserService/entity/implementation"
	"github.com/stretchr/testify/assert"
)

func TestNewPlot(t *testing.T) {
	t.Run("create new plot", func(t *testing.T) {
		plot := entity_impl.NewPlot(1, 1)
		assert.NotNil(t, plot, "instance incorrect")
		assert.Equal(t, "1:1", plot.Id(), "id is incorrect")
	})
}

func TestPlotGetTreeHeight(t *testing.T) {
	t.Run("get tree height after creation", func(t *testing.T) {
		plot := entity_impl.NewPlot(1, 1)
		assert.Equal(t, plot.GetTreeHeight(), 0, "initial height is incorrect")
	})
}

func TestPlotGetSetTreeHeight(t *testing.T) {
	testCases := []struct {
		name           string
		height         int
		expectedHeight int
		expectedError  error
	}{
		{name: "height is less than 1 ", height: -1, expectedHeight: 0, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidTreeHeight)},
		{name: "height is in range (1-30)", height: 13, expectedHeight: 13, expectedError: nil},
		{name: "height is more than 30", height: 671, expectedHeight: 0, expectedError: entity_impl.NewErr(entity_impl.ErrInvalidTreeHeight)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			plot := entity_impl.NewPlot(1, 1)
			err := plot.SetTreeHeight(tc.height)
			height := plot.GetTreeHeight()

			assert.Equal(t, tc.expectedHeight, height, "plot height is incorrect")
			assert.Equal(t, tc.expectedError, err, "error message is incorrect")
		})
	}
}

func TestPlotCoordinate(t *testing.T) {
	testCases := []struct {
		name string
		x, y int
	}{
		{name: "x = 1, y = 1", x: 1, y: 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			plot := entity_impl.NewPlot(tc.x, tc.y)
			x, y := plot.Coordinate()

			assert.Equal(t, tc.x, x, "x is incorrect")
			assert.Equal(t, tc.y, y, "y is incorrect")
		})
	}
}
