package entity_impl

import (
	"github.com/SawitProRecruitment/UserService/entity"
)

var _ entity.Navigator = (*Navigator)(nil)

func NewNavigator(estate entity.Estate, path entity.PathProvider, drone entity.Drone) (*Navigator, error) {
	if estate == nil {
		return nil, NewErr(ErrEstateIsNil)
	}
	if path == nil {
		return nil, NewErr(ErrPathProviderIsNil)
	}
	if drone == nil {
		return nil, NewErr(ErrDroneIsNil)
	}
	return &Navigator{
		started: false,
		estate:  estate,
		path:    path,
		drone:   drone,
	}, nil
}

type Navigator struct {
	started     bool
	estate      entity.Estate
	path        entity.PathProvider
	drone       entity.Drone
	currentPlot entity.Plot
}

func (n *Navigator) Start() error {
	if n.started {
		return NewErr(ErrNavigatorAlreadyStarted)
	}
	n.started = true
	// get start coordinate
	x, y := n.path.Start()

	// get the plot from coordinate
	plot, _ := n.estate.GetPlot(x, y)
	n.currentPlot = plot

	// move the drone
	n.drone.Activate()
	n.drone.Elevate(MIN_ELEVATION_FROM_SURFACE)

	return nil
}

func (n *Navigator) Move(nextPlot entity.Plot) {
	// setup min height for next plot
	n.drone.Elevate(n.currentPlot.GetTreeHeight() + MIN_ELEVATION_FROM_SURFACE - n.drone.Height()).
		// move to next plot
		NextPlot(nextPlot.Coordinate()).
		// adjust to next plot as local plot
		Elevate(nextPlot.GetTreeHeight() + MIN_ELEVATION_FROM_SURFACE - n.drone.Height())
	n.currentPlot = nextPlot
}

func (n *Navigator) Rest() error {
	if !n.started {
		return NewErr(ErrNavigatorNotStarted)
	}
	n.drone.Stop()
	n.started = false
	return nil
}

func (n *Navigator) End() error {
	if err := n.Rest(); err != nil {
		return err
	}
	n.currentPlot = nil
	return nil
}
