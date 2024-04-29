package entity_impl

import (
	"github.com/SawitProRecruitment/UserService/entity"
)

var _ entity.PathProvider = (*PathProvider)(nil)

func NewPathProvider(estate entity.Estate) (*PathProvider, error) {
	if estate == nil {
		return nil, NewErr(ErrEstateIsNil)
	}
	return &PathProvider{
		estate: estate,
	}, nil
}

type PathProvider struct {
	estate entity.Estate
}

func (p *PathProvider) Start() (x int, y int) {
	return 1, 1
}

func (p *PathProvider) Next(currentX, currentY int) (x int, y int) {
	if p.IsEnd(currentX, currentY) {
		return p.Start()
	}
	if currentY%2 == 1 {
		if currentX == p.estate.Length() {
			x = currentX
			y = currentY + 1
		} else {
			x = currentX + 1
			y = currentY
		}
	} else {
		if currentX == 1 {
			x = currentX
			y = currentY + 1
		} else {
			x = currentX - 1
			y = currentY
		}
	}
	return
}

func (p *PathProvider) IsEnd(x int, y int) bool {
	return x == p.estate.Length() && y == p.estate.Width()
}

func (p *PathProvider) IsStart(x int, y int) bool {
	return x == 1 && y == 1
}
