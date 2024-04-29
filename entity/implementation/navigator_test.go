package entity_impl_test

import (
	"testing"

	entity_impl "github.com/SawitProRecruitment/UserService/entity/implementation"
	"github.com/stretchr/testify/assert"
)

func TestNewNavigator(t *testing.T) {
	estate, _ := entity_impl.NewEstate(4, 3)
	pathprov, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(-1, nil)

	nav, err := entity_impl.NewNavigator(nil, pathprov, drone)
	if assert.Error(t, err) {
		assert.Equal(t, entity_impl.NewErr(entity_impl.ErrEstateIsNil), err)
		assert.Nil(t, nav)
	}

	nav, err = entity_impl.NewNavigator(estate, nil, drone)
	if assert.Error(t, err) {
		assert.Equal(t, entity_impl.NewErr(entity_impl.ErrPathProviderIsNil), err)
		assert.Nil(t, nav)
	}

	nav, err = entity_impl.NewNavigator(estate, pathprov, nil)
	if assert.Error(t, err) {
		assert.Equal(t, entity_impl.NewErr(entity_impl.ErrDroneIsNil), err)
		assert.Nil(t, nav)
	}

	nav, err = entity_impl.NewNavigator(estate, pathprov, drone)
	assert.NotNil(t, nav)
	assert.Nil(t, err)
}

func TestNavigatorStart(t *testing.T) {
	estate, _ := entity_impl.NewEstate(4, 3)
	pathprov, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(-1, nil)

	nav, _ := entity_impl.NewNavigator(estate, pathprov, drone)
	nav.Start()
	assert.True(t, drone.Active())
	assert.Equal(t, entity_impl.MIN_ELEVATION_FROM_SURFACE, drone.Height())

	err := nav.Start()
	assert.Equal(t, entity_impl.NewErr(entity_impl.ErrNavigatorAlreadyStarted), err)
}

func TestNavigatorRest(t *testing.T) {
	estate, _ := entity_impl.NewEstate(4, 3)
	pathprov, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(-1, nil)

	nav, _ := entity_impl.NewNavigator(estate, pathprov, drone)
	err := nav.Rest()
	assert.Equal(t, entity_impl.NewErr(entity_impl.ErrNavigatorNotStarted), err)

	err = nav.End()
	assert.Equal(t, entity_impl.NewErr(entity_impl.ErrNavigatorNotStarted), err)

	nav.Start()
	err = nav.End()
	assert.Nil(t, err)
}

func TestNavigatorMove(t *testing.T) {
	testCases := []struct {
		name   string
		x, y   int
		height int
	}{
		{name: "Plot(2, 1); Height: 5", x: 2, y: 1, height: 5},
		{name: "Plot(3, 1); Height: 3", x: 3, y: 1, height: 3},
		{name: "Plot(4, 1); Height: 4", x: 4, y: 1, height: 4},
	}

	estate, _ := entity_impl.NewEstate(1, 5)
	pathprov, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(-1, nil)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			estate.SetTreeHeight(tc.x, tc.y, tc.height)
		})
	}

	nav, _ := entity_impl.NewNavigator(estate, pathprov, drone)
	nav.Start()
	for i := 2; i <= 5; i++ {
		plot, _ := estate.GetPlot(i, 1)
		nav.Move(plot)
	}
	nav.End()
	assert.Equal(t, 54, drone.TotalDistance(), "drone total distance is incorrect")
	assert.Equal(t, 14, drone.VerticalDistance(), "drone vertical distance is incorrect")
	assert.Equal(t, 40, drone.HorizontalDistance(), "drone horizontal distance is incorrect")
}
