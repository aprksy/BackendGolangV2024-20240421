package entity_impl

import (
	"github.com/SawitProRecruitment/UserService/entity"
)

var _ entity.Drone = (*Drone)(nil)

func NewDrone(maxDistance int, onMaxDistanceReached func(x, y, distance int)) (*Drone, error) {
	if maxDistance == 0 || maxDistance < -1 {
		return nil, NewErr(ErrDroneInvalidMaxDistance)
	}
	return &Drone{
		maxDistance:                  maxDistance,
		onMaxDistanceReachedCallback: onMaxDistanceReached,
	}, nil
}

type Drone struct {
	maxDistance                  int
	x, y                         int
	height                       int
	verticalDistance             int
	horizontalDistance           int
	onMaxDistanceReachedCallback func(x, y, distance int)
	active                       bool
}

func (d *Drone) MaxDistance() int {
	return d.maxDistance
}

func (d *Drone) Activate() {
	d.active = true
}

func (d *Drone) Deactivate() {
	d.active = false
}

func (d *Drone) Active() bool {
	return d.active
}

func (d *Drone) Position() (x int, y int) {
	return d.x, d.y
}

func (d *Drone) Height() int {
	return d.height
}

func (d *Drone) VerticalDistance() int {
	return d.verticalDistance
}

func (d *Drone) HorizontalDistance() int {
	return d.horizontalDistance
}

func (d *Drone) TotalDistance() int {
	return d.horizontalDistance + d.verticalDistance
}

func (d *Drone) Elevate(meters int) entity.Drone {
	if !d.active {
		return d
	}

	// check whether remaining distance for safe landing
	// need to choose to land in current plot or next plot
	// if our total distance plus height less then max distance,
	// but max distance is less then current position to the ground
	// on next plot, it will be safer to land here, otherwise we will
	// crash on next plot

	// deltaHeight := 0
	// if meters > 0 {
	// 	deltaHeight = 2 * meters
	// } else {
	// 	deltaHeight = meters
	// }
	// if d.TotalDistance()+d.height <= d.maxDistance &&
	// 	d.maxDistance < d.TotalDistance()+d.height+deltaHeight {
	// 	return d.Stop()
	// } else if d.maxDistance == d.TotalDistance()+d.height+deltaHeight {
	// 	d.verticalDistance = d.verticalDistance + deltaHeight
	// 	return d.Stop()
	// }

	if d.maxDistance != -1 && d.TotalDistance()+meters >= d.maxDistance {
		d.verticalDistance = d.maxDistance - d.horizontalDistance
		d.height = 0
		return d.Stop()
	}

	d.verticalDistance = d.verticalDistance + abs(meters)
	d.height = d.height + meters

	// fmt.Printf("elevate: %d; vdist: %d; hdist: %d; height: %d\n", meters, d.verticalDistance, d.horizontalDistance, d.height)
	return d
}

func (d *Drone) NextPlot(x, y int) entity.Drone {
	if !d.active {
		return d
	}
	// if d.TotalDistance()+d.height <= d.maxDistance &&
	// 	d.maxDistance < d.TotalDistance()+d.height+PLOT_SIZE {
	// 	return d.Stop()
	// } else if d.maxDistance == d.TotalDistance()+d.height+PLOT_SIZE {
	// 	d.horizontalDistance = d.horizontalDistance + PLOT_SIZE
	// 	return d.Stop()
	// }

	d.x = x
	d.y = y

	if d.maxDistance != -1 && d.TotalDistance()+PLOT_SIZE >= d.maxDistance {
		d.horizontalDistance = d.maxDistance - d.verticalDistance
		d.height = 0
		return d.Stop()
	}

	d.horizontalDistance = d.horizontalDistance + PLOT_SIZE
	return d
}

func (d *Drone) Stop() entity.Drone {
	d.verticalDistance = d.verticalDistance + d.height
	d.active = false
	if d.onMaxDistanceReachedCallback != nil {
		d.onMaxDistanceReachedCallback(d.x, d.y, d.TotalDistance())
	}
	return d
}

func (d *Drone) SetOnMaxDistanceReachedCallback(fn func(x, y, distance int)) {
	d.onMaxDistanceReachedCallback = fn
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
