package entity

type MovementDirection int

const (
	DirectionVertical MovementDirection = iota
	DirectionHorizontal
	DirectionNone
)

type Movement struct {
	Distance  int
	Direction MovementDirection
}
