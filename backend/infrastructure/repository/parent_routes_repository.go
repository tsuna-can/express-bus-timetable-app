package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/factory"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

const (
	getAllParentRoutesQuery = "SELECT parent_route_id, parent_route_name FROM parent_route"
	getParentRouteByIdQuery = "SELECT parent_route_id, parent_route_name FROM parent_route WHERE parent_route_id = $1"
)

type ParentRoutesRepository struct {
	db      *sqlx.DB
	factory factory.ParentRouteFactory
}

func NewParentRoutesRepository(db *sqlx.DB) repository.ParentRoutesRepository {
	return &ParentRoutesRepository{
		db:      db,
		factory: factory.NewParentRouteFactory(),
	}
}

func (r *ParentRoutesRepository) GetAll(ctx context.Context) ([]entity.ParentRoute, error) {
	rows, err := r.db.QueryContext(ctx, getAllParentRoutesQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query parent routes: %w", err)
	}
	defer rows.Close()

	var rawDataList []factory.ParentRouteRawData
	for rows.Next() {
		var rawData factory.ParentRouteRawData
		if err := rows.Scan(&rawData.ParentRouteId, &rawData.ParentRouteName); err != nil {
			return nil, fmt.Errorf("failed to scan parent route row: %w", err)
		}
		rawDataList = append(rawDataList, rawData)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	// Convert raw data to entities using factory
	parentRoutes, err := r.factory.ReconstructManyFromRawData(rawDataList)
	if err != nil {
		return nil, fmt.Errorf("failed to reconstruct parent routes from raw data: %w", err)
	}

	return parentRoutes, nil
}

func (r *ParentRoutesRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) (entity.ParentRoute, error) {
	var rawData factory.ParentRouteRawData

	row := r.db.QueryRowContext(ctx, getParentRouteByIdQuery, parentRouteId)
	if err := row.Scan(&rawData.ParentRouteId, &rawData.ParentRouteName); err != nil {
		return entity.ParentRoute{}, fmt.Errorf("failed to scan parent route: %w", err)
	}

	parentRoute, err := r.factory.ReconstructFromRawData(rawData)
	if err != nil {
		return entity.ParentRoute{}, fmt.Errorf("failed to reconstruct parent route from raw data: %w", err)
	}

	return *parentRoute, nil
}
