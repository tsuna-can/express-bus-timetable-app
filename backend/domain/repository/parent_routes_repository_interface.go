package repository

import (
  "context"
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type ParentRoutesRepository interface {
  GetAll(ctx context.Context) ([]entity.ParentRoute, error)
}

