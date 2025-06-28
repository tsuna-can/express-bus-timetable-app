package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/model"
)

const (
	getAllParentRoutesQuery = "SELECT parent_route_id, parent_route_name FROM parent_route"
	getParentRouteByIdQuery = "SELECT parent_route_id, parent_route_name FROM parent_route WHERE parent_route_id = $1"
)

type ParentRoutesRepository struct {
	db *sqlx.DB
}

func NewParentRoutesRepository(db *sqlx.DB) repository.ParentRoutesRepository {
	return &ParentRoutesRepository{
		db: db,
	}
}

func (r *ParentRoutesRepository) GetAll(ctx context.Context) ([]entity.ParentRoute, error) {
	rows, err := r.db.QueryContext(ctx, getAllParentRoutesQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query parent routes: %w", err)
	}
	defer rows.Close()

	var parentRoutes []entity.ParentRoute
	for rows.Next() {
		var prm model.ParentRoute
		if err := rows.Scan(&prm.ParentRouteId, &prm.ParentRouteName); err != nil {
			return nil, fmt.Errorf("failed to scan parent route row: %w", err)
		}

		pre, err := prm.ToParentRoute()
		if err != nil {
			return nil, fmt.Errorf("failed to convert to ParentRoute: %w", err)
		}

		parentRoutes = append(parentRoutes, *pre)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	return parentRoutes, nil
}

func (r *ParentRoutesRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) (entity.ParentRoute, error) {
	var prm model.ParentRoute

	row := r.db.QueryRowContext(ctx, getParentRouteByIdQuery, parentRouteId)
	if err := row.Scan(&prm.ParentRouteId, &prm.ParentRouteName); err != nil {
		return entity.ParentRoute{}, fmt.Errorf("failed to scan parent route: %w", err)
	}

	pre, err := prm.ToParentRoute()
	if err != nil {
		return entity.ParentRoute{}, fmt.Errorf("failed to convert to ParentRoute: %w", err)
	}

	return *pre, nil
}
