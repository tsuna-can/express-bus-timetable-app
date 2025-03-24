package repository

import (
  "context"
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type TimetableRepository interface {
  GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error)
}

