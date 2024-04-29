package repository

import (
	"context"
	"database/sql"

	"github.com/SawitProRecruitment/UserService/entity"
	entity_impl "github.com/SawitProRecruitment/UserService/entity/implementation"
	"github.com/google/uuid"
)

func (r *Repository) CreateEstate(ctx context.Context, input CreateEstateInput) (output CreateEstateOutput, err error) {
	// Prepare estate insertion
	estateStmt, err := r.Db.Prepare("INSERT INTO estate (id, length, width) VALUES ($1, $2, $3)")
	if err != nil {
		return output, err
	}
	defer estateStmt.Close()

	// Execute estate insertion with actual values
	id := uuid.New()
	_, err = estateStmt.Exec(id, input.Length, input.Width)
	if err != nil {
		return output, err
	}
	output.Id = id
	return output, nil
}

func (r *Repository) SetTreeHeight(ctx context.Context, input SetTreeHeightInput) (output SetTreeHeightOutput, err error) {
	// Prepare plot insertion
	plotStmt, err := r.Db.Prepare("INSERT INTO plot (estate_id, x, y, height) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return output, err
	}
	defer plotStmt.Close()

	// Execute plot insertion with actual values (assuming estateID is used)
	_, err = plotStmt.Exec(input.Id, input.X, input.Y, input.Height)
	if err != nil {
		return output, err
	}
	output.Id = input.Id
	return output, nil
}

func (r *Repository) GetEstateStats(ctx context.Context, input GetEstateStatsInput) (output GetEstateStatsOutput, err error) {
	output = GetEstateStatsOutput{
		Count:  0,
		Min:    0,
		Max:    0,
		Median: 0,
	}
	// Prepare stats query
	statsStmt, err := r.Db.Prepare("SELECT * FROM estate_stats WHERE estate_id = $1::uuid;")
	if err != nil {
		return output, err
	}
	defer statsStmt.Close()

	// Execute stats query
	row := statsStmt.QueryRow(input.Id)
	var id uuid.UUID
	err = row.Scan(&id, &output.Count, &output.Min, &output.Max, &output.Median)
	if err != nil {
		if err == sql.ErrNoRows {
			return output, nil
		}
		return output, err
	}
	return output, nil
}

func (r *Repository) PlanDronePatrol(ctx context.Context, input PlanDronePatrolInput) (output PlanDronePatrolOutput, err error) {
	stmt, err := r.Db.Prepare(`SELECT e.id AS id, e.length as length, e.width as width, p.x as x, p.y as y, p.height as height
	FROM estate e
	INNER JOIN plot p ON p.estate_id = e.id
	WHERE e.id = $1::uuid;`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(input.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return output, nil
		}
		return output, err
	}
	defer rows.Close()

	var (
		estate entity.Estate
	)
	for rows.Next() {
		var (
			id           uuid.UUID
			length       int
			width        int
			x, y, height int
		)
		err := rows.Scan(&id, &length, &width, &x, &y, &height)
		if err != nil {
			return output, err
		}
		if estate == nil {
			estate, _ = entity_impl.NewEstate(width, length)
		}
		err = estate.SetTreeHeight(x, y, height)
		if err != nil {
			return output, err
		}
	}

	pathProv, _ := entity_impl.NewPathProvider(estate)
	drone, _ := entity_impl.NewDrone(input.MaxDistance, nil)
	nav, _ := entity_impl.NewNavigator(estate, pathProv, drone)
	patrol, _ := entity_impl.NewPatrol(estate, pathProv, nav, drone)
	drone.SetOnMaxDistanceReachedCallback(patrol.RecordRestPlot)

	rec := patrol.Plan()

	output.Distance = rec.Distance
	output.Rest.X = rec.Rest.X
	output.Rest.Y = rec.Rest.Y
	return output, nil
}
