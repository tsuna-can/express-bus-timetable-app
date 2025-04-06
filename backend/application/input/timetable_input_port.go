package input

import (
	"context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type TimetableInputPort interface {
	GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error)
}
