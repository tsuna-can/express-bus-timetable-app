package repository

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/model"
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
	var parentRoutes []entity.ParentRoute

	rows, err := r.db.QueryContext(ctx, "SELECT parent_route_id, parent_route_name FROM parent_route")
	if err != nil {
		log.Printf("Error querying parent routes: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var prm model.ParentRoute
		if err := rows.Scan(&prm.ParentRouteId, &prm.ParentRouteName); err != nil {
			log.Printf("Error scanning parent route: %v", err)
			return nil, err
		}

		pre, err := prm.ToParentRoute()
		if err != nil {
			log.Printf("Error converting to ParentRoute: %v", err)
			return nil, err
		}

		parentRoutes = append(parentRoutes, *pre)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating parent routes: %v", err)
		return nil, err
	}

	return parentRoutes, nil
}

func (r *ParentRoutesRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) (entity.ParentRoute, error) {
	var (
		parentRouteIdRaw   string
		parentRouteNameRaw string
	)

	row := r.db.QueryRowContext(ctx, "SELECT parent_route_id, parent_route_name FROM parent_route WHERE parent_route_id = $1", parentRouteId)
	if err := row.Scan(&parentRouteIdRaw, &parentRouteNameRaw); err != nil {
		log.Printf("Error scanning parent route: %v", err)
		return entity.ParentRoute{}, err
	}

	parentRouteName, prnErr := vo.NewParentRouteName(parentRouteNameRaw)
	if prnErr != nil {
		log.Printf("Error creating ParentRouteName: %v", prnErr)
		return entity.ParentRoute{}, prnErr
	}

	parentRoute := entity.ParentRoute{
		ParentRouteId:   parentRouteId,
		ParentRouteName: *parentRouteName,
	}

	return parentRoute, nil
}
