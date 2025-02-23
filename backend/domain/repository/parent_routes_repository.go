package repository

import (
  "context"
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/model"
)

type ParentRouteRepository interface {
  GetAll(ctx context.Context) ([]model.ParentRoute, error)
}

