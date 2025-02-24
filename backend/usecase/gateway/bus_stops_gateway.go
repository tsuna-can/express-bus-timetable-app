package gateway

import (
  "context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type BusStopsGateway interface {
	GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, error)
}

