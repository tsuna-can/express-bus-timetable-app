package gateway

import (
  "context"
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type TimetableGateway interface {
  GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error)
}

