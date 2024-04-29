package entity_impl

import (
	"github.com/SawitProRecruitment/UserService/entity"
)

func NewPatrol(estate entity.Estate, pathProv entity.PathProvider,
	navigator entity.Navigator, drone entity.Drone) (*Patrol, error) {
	if estate == nil {
		return nil, NewErr(ErrEstateIsNil)
	}
	if pathProv == nil {
		return nil, NewErr(ErrPathProviderIsNil)
	}
	if navigator == nil {
		return nil, NewErr(ErrNavigatorIsNil)
	}
	if drone == nil {
		return nil, NewErr(ErrDroneIsNil)
	}
	return &Patrol{
		estate:       estate,
		pathProvider: pathProv,
		navigator:    navigator,
		drone:        drone,
		record: PatrolRecord{
			Distance: 0,
			Rest:     PatrolRest{},
		},
	}, nil
}

type PatrolRest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PatrolRecord struct {
	Distance int        `json:"distance"`
	Rest     PatrolRest `json:"rest"`
}

type Patrol struct {
	estate       entity.Estate
	pathProvider entity.PathProvider
	navigator    entity.Navigator
	drone        entity.Drone
	record       PatrolRecord
}

func (p *Patrol) Plan() PatrolRecord {
	p.navigator.Start()

	x, y := 1, 1

	for !p.pathProvider.IsEnd(x, y) && p.drone.Active() {
		x, y = p.pathProvider.Next(x, y)
		plot, _ := p.estate.GetPlot(x, y)
		p.navigator.Move(plot)
	}
	p.navigator.End()

	return p.record
}

func (p *Patrol) RecordRestPlot(x, y, distance int) {
	rest := PatrolRest{
		X: x,
		Y: y,
	}
	p.record.Distance = distance
	p.record.Rest = rest
}
