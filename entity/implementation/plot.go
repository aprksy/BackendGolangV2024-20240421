package entity_impl

import "fmt"

func NewPlot(x, y int) *Plot {
	return &Plot{
		x:      x,
		y:      y,
		height: 0,
	}
}

type Plot struct {
	x, y   int
	height int
}

func (p *Plot) Id() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func (p *Plot) Coordinate() (x, y int) {
	return p.x, p.y
}

func (p *Plot) GetTreeHeight() int {
	return p.height
}

func (p *Plot) SetTreeHeight(value int) error {
	if value < PLOT_MIN_TREE_HEIGHT || value > PLOT_MAX_TREE_HEIGHT {
		return NewErr(ErrInvalidTreeHeight)
	}
	p.height = value
	return nil
}
