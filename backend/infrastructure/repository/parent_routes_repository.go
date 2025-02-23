package repository

import (
  "log"
  "context"
	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/model"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)


type ParentRouteRepository struct {
	db *sqlx.DB
}

func NewParentRouteRepository(db *sqlx.DB) repository.ParentRouteRepository {
	return &ParentRouteRepository{
		db: db,
	}
}

func (r *ParentRouteRepository) GetAll(ctx context.Context) ([]model.ParentRoute, error) {
  var parentRoutes []model.ParentRoute

  rows, err := r.db.QueryContext(ctx, "SELECT parent_route_id, parent_route_name FROM parentroute")
  if err != nil {
    log.Printf("Error querying parent routes: %v", err)
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var parentRoute model.ParentRoute
    if err := rows.Scan(&parentRoute.ParentRouteId, &parentRoute.ParentRouteName); err != nil {
      log.Printf("Error scanning parent route: %v", err)
      return nil, err
    }
    parentRoutes = append(parentRoutes, parentRoute)
  }

  if err := rows.Err(); err != nil {
    log.Printf("Error iterating parent routes: %v", err)
    return nil, err
  }

  return parentRoutes, nil
}

