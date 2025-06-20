package usecase 

import (
	"context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/application/input"
	"log"
)

type timetableUsecase struct {
	timetableRepository repository.TimetableRepository
}

func NewTimetableUsecase(timetableRepository repository.TimetableRepository) input.TimetableInputPort {
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
