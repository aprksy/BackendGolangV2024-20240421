package entity_impl

import (
	"github.com/SawitProRecruitment/UserService/entity"
	"github.com/google/uuid"
)

func NewEstate(width, length int) (*Estate, error) {
	if width < ESTATE_MIN_WIDTH || width > ESTATE_MAX_WIDTH ||
		length < ESTATE_MIN_LENGTH || length > ESTATE_MAX_LENGTH {
		return nil, NewErr(ErrInvalidEstateDimension)
	}

	plots := [][]entity.Plot{}
	for i := 0; i < width; i++ {
		plots = append(plots, []entity.Plot{})
		for j := 0; j < length; j++ {
			plots[i] = append(plots[i], NewPlot(j+1, i+1))
		}
	}

	return &Estate{
		id:     uuid.New(),
		width:  width,
		length: length,
		plots:  plots,
	}, nil
}

type Estate struct {
	id     uuid.UUID
	width  int
	length int
	plots  [][]entity.Plot
}

func (e *Estate) Id() uuid.UUID {
	return e.id
}

func (e *Estate) Length() int {
	return e.length
}

func (e *Estate) Width() int {
	return e.width
}

func (e *Estate) SetTreeHeight(x int, y int, value int) error {
	plot, err := e.GetPlot(x, y)
	if err != nil {
		return err
	}
	if err := plot.SetTreeHeight(value); err != nil {
		return err
	}
	return nil
}

func (e *Estate) GetTreeHeight(x int, y int) (height int, err error) {
	plot, err := e.GetPlot(x, y)
	if err != nil {
		return -1, err
	}
	return plot.GetTreeHeight(), nil
}

func (e *Estate) GetPlot(x int, y int) (plot entity.Plot, err error) {
	if x < 1 || x > e.length ||
		y < 1 || y > e.width {
		return nil, NewErr(ErrInvalidCoordinates)
	}
	return e.plots[y-1][x-1], nil
}
