package entity_impl_test

import (
	"testing"

	entity_impl "github.com/SawitProRecruitment/UserService/entity/implementation"
	"github.com/stretchr/testify/assert"
)

func TestNewDrone(t *testing.T) {
	testCases := []struct {
		name        string
		maxDistance int
		expectErr   error
	}{
		// failure cases
		{name: "max distance = -3", maxDistance: -3},
		{name: "max distance = -2", maxDistance: -2},
		{name: "max distance = 0", maxDistance: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			drone, err := entity_impl.NewDrone(tc.maxDistance, nil)
			if assert.Error(t, err) {
				assert.Equal(t, entity_impl.NewErr(entity_impl.ErrDroneInvalidMaxDistance), err, "err is incorrect")
				assert.Nil(t, drone, "drone instance should be nil")
			}
		})
	}

	testCases1 := []struct {
		name        string
		maxDistance int
	}{
		// success cases
		{name: "max distance = -1", maxDistance: -1},
		{name: "max distance = 2", maxDistance: 2},
		{name: "max distance = 10", maxDistance: 10},
	}

	for _, tc := range testCases1 {
		t.Run(tc.name, func(t *testing.T) {
			drone, err := entity_impl.NewDrone(tc.maxDistance, nil)
			assert.NotNil(t, drone, "drone instance should not be nil")
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, tc.maxDistance, drone.MaxDistance(), "drone max distance is incorrect")
		})
	}
}

func TestDroneActivateDeactivate(t *testing.T) {
	drone, _ := entity_impl.NewDrone(100, nil)

	drone.Activate()
	assert.True(t, drone.Active())

	drone.Deactivate()
	assert.False(t, drone.Active())
}

func sampleCallback(x, y, distance int) {

}

func TestDroneElevate(t *testing.T) {
	testCases1 := []struct {
		name       string
		moveMeters int
		height     int
		distance   int
	}{
		// success cases
		{name: "elevate 1", moveMeters: 1, distance: 1, height: 1},
		{name: "elevate 8", moveMeters: 8, distance: 9, height: 9},
		{name: "elevate 11", moveMeters: 11, distance: 20, height: 20},
		{name: "elevate -5", moveMeters: -5, distance: 25, height: 15},
		{name: "elevate -4", moveMeters: -4, distance: 29, height: 11},
		{name: "elevate 15", moveMeters: 15, distance: 44, height: 26},
		{name: "elevate -20", moveMeters: -20, distance: 64, height: 6},
		{name: "elevate -6", moveMeters: -6, distance: 70, height: 0},
	}

	drone, _ := entity_impl.NewDrone(-1, sampleCallback)
	drone.Elevate(100)
	assert.Equal(t, 0, drone.VerticalDistance(), "drone vertical distance is incorrect")

	drone.Activate()
	for _, tc := range testCases1 {
		t.Run(tc.name, func(t *testing.T) {
			drone.Elevate(tc.moveMeters)
			assert.Equal(t, tc.distance, drone.VerticalDistance(), "drone vertical distance is incorrect")
			assert.Equal(t, 0, drone.HorizontalDistance(), "drone horizontal distance is incorrect")
			assert.Equal(t, tc.height, drone.Height(), "drone height is incorrect")
			assert.Equal(t, tc.distance, drone.TotalDistance(), "drone total distance is incorrect")
		})
	}

	drone.Stop()
	assert.Equal(t, 0, drone.Height(), "drone height is incorrect")

	// test for drone that need to rest 1
	drone, _ = entity_impl.NewDrone(50, nil)
	drone.Activate()
	// success
	drone.Elevate(10)
	assert.Equal(t, 10, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 10, drone.Height(), "drone height is incorrect")
	// success
	drone.Elevate(5)
	assert.Equal(t, 15, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 15, drone.Height(), "drone height is incorrect")
	// success, but need to rest immediately otherwise it will crash on next plot
	drone.Elevate(35)
	assert.Equal(t, 50, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 50, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 0, drone.Height(), "drone height is incorrect")

	// test for drone that need to rest 2
	drone, _ = entity_impl.NewDrone(50, nil)
	drone.Activate()
	// success
	drone.Elevate(10)
	assert.Equal(t, 10, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 10, drone.Height(), "drone height is incorrect")
	// can elevate, but choose to rest because it will crash on next plot
	drone.Elevate(60)
	assert.Equal(t, 50, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 50, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 0, drone.Height(), "drone height is incorrect")
}

func TestDroneNextPlot(t *testing.T) {
	testCases1 := []struct {
		name     string
		x        int
		y        int
		distance int
	}{
		// success cases
		{name: "elevate 1", x: 1, y: 1, distance: 1 * entity_impl.PLOT_SIZE},
		{name: "elevate 2", x: 2, y: 7, distance: 2 * entity_impl.PLOT_SIZE},
		{name: "elevate 3", x: 3, y: 5, distance: 3 * entity_impl.PLOT_SIZE},
		{name: "elevate 4", x: 5, y: 2, distance: 4 * entity_impl.PLOT_SIZE},
		{name: "elevate 5", x: 8, y: 7, distance: 5 * entity_impl.PLOT_SIZE},
		{name: "elevate 6", x: 2, y: 9, distance: 6 * entity_impl.PLOT_SIZE},
		{name: "elevate 7", x: 4, y: 5, distance: 7 * entity_impl.PLOT_SIZE},
	}

	drone, _ := entity_impl.NewDrone(-1, sampleCallback)
	drone.NextPlot(100, 3)
	assert.Equal(t, 0, drone.VerticalDistance(), "drone horizontal distance is incorrect")

	drone, _ = entity_impl.NewDrone(-1, sampleCallback)
	drone.Activate()

	for _, tc := range testCases1 {
		t.Run(tc.name, func(t *testing.T) {
			drone.NextPlot(tc.x, tc.y)
			x, y := drone.Position()
			assert.Equal(t, 0, drone.VerticalDistance(), "drone vertical distance is incorrect")
			assert.Equal(t, tc.distance, drone.HorizontalDistance(), "drone horizontal distance is incorrect")
			assert.Equal(t, tc.distance, drone.TotalDistance(), "drone total distance is incorrect")
			assert.Equal(t, tc.x, x, "drone x position is incorrect")
			assert.Equal(t, tc.y, y, "drone y position is incorrect")
		})
	}

	// test for drone that need to rest 1
	drone, _ = entity_impl.NewDrone(30, nil)
	drone.Activate()
	// success
	drone.Elevate(10)
	assert.Equal(t, 10, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 10, drone.Height(), "drone height is incorrect")
	// success
	drone.NextPlot(2, 1)
	assert.Equal(t, 20, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 10, drone.Height(), "drone height is incorrect")
	// success, but need to rest immediately otherwise it will crash on next plot
	drone.NextPlot(3, 1)
	assert.Equal(t, 30, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 0, drone.Height(), "drone height is incorrect")

	// test for drone that need to rest 2
	drone, _ = entity_impl.NewDrone(35, nil)
	drone.Activate()
	// success, elevating to 23
	drone.Elevate(23)
	assert.Equal(t, 23, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 23, drone.Height(), "drone height is incorrect")
	// success, go to next plot
	drone.NextPlot(2, 1)
	assert.Equal(t, 33, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 23, drone.Height(), "drone height is incorrect")
	// can't continue
	drone.NextPlot(2, 1)
	assert.Equal(t, 35, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 0, drone.Height(), "drone height is incorrect")
}
