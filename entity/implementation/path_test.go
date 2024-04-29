package entity_impl_test

import (
	"testing"

	entity_impl "github.com/SawitProRecruitment/UserService/entity/implementation"
	"github.com/stretchr/testify/assert"
)

func TestNewPathProvider(t *testing.T) {
	t.Run("create new path", func(t *testing.T) {
		path, err := entity_impl.NewPathProvider(nil)
		if assert.Error(t, err) {
			assert.Equal(t, entity_impl.NewErr(entity_impl.ErrEstateIsNil), err)
			assert.Nil(t, path)
		}

		estate, _ := entity_impl.NewEstate(1, 1)
		path, err = entity_impl.NewPathProvider(estate)
		assert.NotNil(t, path, "instance incorrect")
		assert.Nil(t, err, "err should be nil")
	})
}

func TestPathProviderStart(t *testing.T) {
	t.Run("get the start plot", func(t *testing.T) {
		estate, _ := entity_impl.NewEstate(4, 3)
		path, _ := entity_impl.NewPathProvider(estate)
		x, y := path.Start()
		assert.Equal(t, x, 1, "x should be 1")
		assert.Equal(t, y, 1, "y should be 1")
	})
}

func TestPathProviderIsEnd(t *testing.T) {
	testCases := []struct {
		name   string
		length int
		width  int
		x      int
		y      int
		result bool
	}{
		{name: "estate 1 x 1 ", length: 1, width: 1, x: 1, y: 1, result: true},
		{name: "estate 3 x 4 ", length: 3, width: 4, x: 3, y: 4, result: true},
		{name: "estate 4 x 5 ", length: 4, width: 5, x: 4, y: 5, result: true},
		{name: "estate 2 x 8 ", length: 2, width: 8, x: 2, y: 8, result: true},
		{name: "estate 9 x 1 ", length: 9, width: 1, x: 9, y: 1, result: true},
		{name: "estate 3 x 3 ", length: 3, width: 3, x: 3, y: 3, result: true},
		{name: "estate 1 x 7 ", length: 1, width: 7, x: 1, y: 7, result: true},
		{name: "estate 1 x 1 ", length: 1, width: 1, x: 1, y: 2, result: false},
		{name: "estate 3 x 4 ", length: 3, width: 4, x: 3, y: 2, result: false},
		{name: "estate 4 x 5 ", length: 4, width: 5, x: 4, y: 2, result: false},
		{name: "estate 2 x 8 ", length: 2, width: 8, x: 2, y: 2, result: false},
		{name: "estate 9 x 1 ", length: 9, width: 1, x: 9, y: 2, result: false},
		{name: "estate 3 x 3 ", length: 3, width: 3, x: 3, y: 2, result: false},
		{name: "estate 1 x 7 ", length: 1, width: 7, x: 1, y: 2, result: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, _ := entity_impl.NewEstate(tc.width, tc.length)
			path, _ := entity_impl.NewPathProvider(estate)

			assert.Equal(t, tc.result, path.IsEnd(tc.x, tc.y))
		})
	}
}

func TestPathProviderIsStart(t *testing.T) {
	estate, _ := entity_impl.NewEstate(3, 4)
	path, _ := entity_impl.NewPathProvider(estate)

	assert.True(t, path.IsStart(1, 1))
}

func TestPathProviderNext(t *testing.T) {
	testCases := []struct {
		name   string
		width  int
		length int
		currX  int
		currY  int
		nextX  int
		nextY  int
	}{
		{name: "plot 1", width: 3, length: 4, currX: 1, currY: 1, nextX: 2, nextY: 1},
		{name: "plot 2", width: 3, length: 4, currX: 2, currY: 1, nextX: 3, nextY: 1},
		{name: "plot 3", width: 3, length: 4, currX: 3, currY: 1, nextX: 4, nextY: 1},
		{name: "plot 4", width: 3, length: 4, currX: 4, currY: 1, nextX: 4, nextY: 2},
		{name: "plot 5", width: 3, length: 4, currX: 4, currY: 2, nextX: 3, nextY: 2},
		{name: "plot 6", width: 3, length: 4, currX: 3, currY: 2, nextX: 2, nextY: 2},
		{name: "plot 7", width: 3, length: 4, currX: 2, currY: 2, nextX: 1, nextY: 2},
		{name: "plot 8", width: 3, length: 4, currX: 1, currY: 2, nextX: 1, nextY: 3},
		{name: "plot 9", width: 3, length: 4, currX: 1, currY: 3, nextX: 2, nextY: 3},
		{name: "plot 10", width: 3, length: 4, currX: 2, currY: 3, nextX: 3, nextY: 3},
		{name: "plot 11", width: 3, length: 4, currX: 3, currY: 3, nextX: 4, nextY: 3},
		{name: "plot 12", width: 3, length: 4, currX: 4, currY: 3, nextX: 1, nextY: 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate, _ := entity_impl.NewEstate(tc.width, tc.length)
			path, _ := entity_impl.NewPathProvider(estate)
			x, y := path.Next(tc.currX, tc.currY)

			assert.Equal(t, tc.nextX, x, "x value is incorrect")
			assert.Equal(t, tc.nextY, y, "y value is incorrect")
		})
	}
}
