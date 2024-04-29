// This file contains types that are used in the repository layer.
package repository

import (
	"github.com/SawitProRecruitment/UserService/entity"
	"github.com/google/uuid"
)

type GetEstateByIdInput struct {
	Id uuid.UUID
}

type GetEstateByIdOutput struct {
	Estate entity.Estate
}

type CreateEstateInput struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}

type CreateEstateOutput struct {
	Id uuid.UUID `json:"id"`
}

type SetTreeHeightInput struct {
	Id     uuid.UUID
	X      int `json:"x"`
	Y      int `json:"y"`
	Height int `json:"height"`
}

type SetTreeHeightOutput struct {
	Id uuid.UUID `json:"id"`
}

type GetEstateStatsInput struct {
	Id uuid.UUID
}

type GetEstateStatsOutput struct {
	Count  int `json:"count"`
	Min    int `json:"min"`
	Max    int `json:"max"`
	Median int `json:"median"`
}

type PlanDronePatrolInput struct {
	Id          uuid.UUID
	MaxDistance int
}

type Rest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlanDronePatrolOutput struct {
	Distance int  `json:"distance"`
	Rest     Rest `json:"rest"`
}
