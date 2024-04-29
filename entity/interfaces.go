package entity

import "github.com/google/uuid"

type Plot interface {
	Id() string
	Coordinate() (x, y int)
	GetTreeHeight() int
	SetTreeHeight(value int) error
}

type Estate interface {
	Id() uuid.UUID
	Length() int
	Width() int
	SetTreeHeight(x, y, value int) error
	GetTreeHeight(x, y int) (height int, err error)
	GetPlot(x, y int) (plot Plot, err error)
}

type Patrol interface {
	Plan() (distance int, rests []Plot)
}

type PathProvider interface {
	Start() (x, y int)
	Next(currentX, currentY int) (x, y int)
	IsEnd(x, y int) bool
	IsStart(x, y int) bool
}

type Navigator interface {
	Start() error
	Move(nextPlot Plot)
	Rest() error
	End() error
}

type Drone interface {
	MaxDistance() int
	Activate()
	Deactivate()
	Active() bool
	Position() (x, y int)
	Height() int
	VerticalDistance() int
	HorizontalDistance() int
	TotalDistance() int
	Elevate(meters int) Drone
	NextPlot(x, y int) Drone
	Stop() Drone
}
