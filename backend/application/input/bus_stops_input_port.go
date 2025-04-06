package input

import (
	"context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type BusStopsInputPort interface {
	GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, error)
}

