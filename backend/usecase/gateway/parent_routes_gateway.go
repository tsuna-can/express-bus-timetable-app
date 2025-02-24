package gateway

import (
  "context"
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type ParentRoutesGateway interface {
  GetAll(ctx context.Context) ([]entity.ParentRoute, error)
}

