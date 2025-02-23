package repository

import (
  "context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type BusStopRepository interface {
	GetByParentRouteId(ctx context.Context, parentRouteId string) ([]model.BusStop, error)
}

