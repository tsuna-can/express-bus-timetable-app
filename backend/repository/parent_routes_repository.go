package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/gateway"
	"log"
  "github.com/tsuna-can/express-bus-time-table-app/backend/repository/model"
)

type ParentRoutesRepository struct {
	db *sqlx.DB
}

func NewParentRoutesRepository(db *sqlx.DB) gateway.ParentRoutesGateway {
	return &ParentRoutesRepository{
		db: db,
	}
}

func (r *ParentRoutesRepository) GetAll(ctx context.Context) ([]entity.ParentRoute, error) {
	var parentRoutes []entity.ParentRoute

	rows, err := r.db.QueryContext(ctx, "SELECT parent_route_id, parent_route_name FROM parentroute")
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
