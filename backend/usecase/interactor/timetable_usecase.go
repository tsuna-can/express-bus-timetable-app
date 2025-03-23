package interactor

import (
	"context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/gateway"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/input_port"
	"log"
)

type timetableUsecase struct {
	timetableRepository gateway.TimetableGateway
}

func NewTimetableUsecase(timetableRepository gateway.TimetableGateway) input_port.TimetableInputPort {
	return &timetableUsecase{
		timetableRepository: timetableRepository,
	}
}

func (u *timetableUsecase) GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error) {
	timetable, err := u.timetableRepository.GetByParentRouteIdAndBusStopId(ctx, parentRouteId, busStopId)
	if err != nil {
		log.Printf("Error getting timetables: %v", err)
    return entity.Timetable{}, err
	}
	return timetable, nil
}
