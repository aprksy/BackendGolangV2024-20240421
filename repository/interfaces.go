// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	CreateEstate(ctx context.Context, input CreateEstateInput) (output CreateEstateOutput, err error)
	SetTreeHeight(ctx context.Context, input SetTreeHeightInput) (output SetTreeHeightOutput, err error)
	GetEstateStats(ctx context.Context, input GetEstateStatsInput) (output GetEstateStatsOutput, err error)
	PlanDronePatrol(ctx context.Context, input PlanDronePatrolInput) (output PlanDronePatrolOutput, err error)
}
