package entity_impl_test

import (
	"testing"

	entity_impl "github.com/SawitProRecruitment/UserService/entity/implementation"
	"github.com/stretchr/testify/assert"
)

func TestNewPatrol(t *testing.T) {
	estate, _ := entity_impl.NewEstate(4, 3)
	pathprov, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(-1, nil)
	nav, _ := entity_impl.NewNavigator(nil, pathprov, drone)

	patrol, err := entity_impl.NewPatrol(nil, pathprov, nav, drone)
	if assert.Error(t, err) {
		assert.Equal(t, entity_impl.NewErr(entity_impl.ErrEstateIsNil), err)
		assert.Nil(t, patrol)
	}

	patrol, err = entity_impl.NewPatrol(estate, nil, nav, drone)
	if assert.Error(t, err) {
		assert.Equal(t, entity_impl.NewErr(entity_impl.ErrPathProviderIsNil), err)
		assert.Nil(t, patrol)
	}

	patrol, err = entity_impl.NewPatrol(estate, pathprov, nil, drone)
	if assert.Error(t, err) {
		assert.Equal(t, entity_impl.NewErr(entity_impl.ErrNavigatorIsNil), err)
		assert.Nil(t, patrol)
	}

	patrol, err = entity_impl.NewPatrol(estate, pathprov, nav, nil)
	if assert.Error(t, err) {
		assert.Equal(t, entity_impl.NewErr(entity_impl.ErrDroneIsNil), err)
		assert.Nil(t, patrol)
	}

	patrol, err = entity_impl.NewPatrol(estate, pathprov, nav, drone)
	assert.NotNil(t, patrol)
	assert.Nil(t, err)
}

func TestPatrolPlan(t *testing.T) {
	estate, _ := entity_impl.NewEstate(1, 5)
	pathprov, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(-1, nil)
	nav, _ := entity_impl.NewNavigator(estate, pathprov, drone)

	testCases := []struct {
		name   string
		x, y   int
		height int
	}{
		{name: "Plot(2, 1); Height: 5", x: 2, y: 1, height: 5},
		{name: "Plot(3, 1); Height: 3", x: 3, y: 1, height: 3},
		{name: "Plot(4, 1); Height: 4", x: 4, y: 1, height: 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate.SetTreeHeight(tc.x, tc.y, tc.height)
		})
	}

	patrol, _ := entity_impl.NewPatrol(estate, pathprov, nav, drone)
	drone.SetOnMaxDistanceReachedCallback(patrol.RecordRestPlot)

	rec := patrol.Plan()
	assert.Equal(t, 54, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 14, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 40, drone.HorizontalDistance(), "drone horizontal distance is incorrect")
	assert.Equal(t, 54, rec.Distance, "distance is incorrect")
	assert.Equal(t, 5, rec.Rest.X, "x is incorrect")
	assert.Equal(t, 1, rec.Rest.Y, "y is incorrect")
}

func TestPatrolPlanWithMaxDistance(t *testing.T) {
	estate, _ := entity_impl.NewEstate(1, 5)
	pathprov, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(25, nil)
	nav, _ := entity_impl.NewNavigator(estate, pathprov, drone)

	testCases := []struct {
		name   string
		x, y   int
		height int
	}{
		{name: "Plot(2, 1); Height: 5", x: 2, y: 1, height: 5},
		{name: "Plot(3, 1); Height: 3", x: 3, y: 1, height: 3},
		{name: "Plot(4, 1); Height: 4", x: 4, y: 1, height: 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate.SetTreeHeight(tc.x, tc.y, tc.height)
		})
	}

	patrol, _ := entity_impl.NewPatrol(estate, pathprov, nav, drone)
	drone.SetOnMaxDistanceReachedCallback(patrol.RecordRestPlot)

	rec := patrol.Plan()
	assert.Equal(t, 25, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 6, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 19, drone.HorizontalDistance(), "drone horizontal distance is incorrect")
	assert.Equal(t, 25, rec.Distance, "distance is incorrect")
	assert.Equal(t, 3, rec.Rest.X, "x is incorrect")
	assert.Equal(t, 1, rec.Rest.Y, "y is incorrect")
}
