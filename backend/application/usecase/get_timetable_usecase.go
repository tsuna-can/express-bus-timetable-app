package usecase

import (
	"context"
	"log"

	"github.com/tsuna-can/express-bus-time-table-app/backend/application/input"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

type getTimetableUsecase struct {
	timetableRepository repository.TimetableRepository
}

func NewGetTimetableUsecase(timetableRepository repository.TimetableRepository) input.TimetableInputPort {
	return &getTimetableUsecase{
		timetableRepository: timetableRepository,
	}
}

func (u *getTimetableUsecase) GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error) {
	timetable, err := u.timetableRepository.GetByParentRouteIdAndBusStopId(ctx, parentRouteId, busStopId)
	if err != nil {
		log.Printf("Error getting timetables: %v", err)
		return entity.Timetable{}, err
	}
	return timetable, nil
}
