package usecase

import (
	"context"
	"fmt"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

type GetTimetableUsecase interface {
	GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error)
}

type getTimetableUsecase struct {
	timetableRepository repository.TimetableRepository
}

func NewGetTimetableUsecase(timetableRepository repository.TimetableRepository) GetTimetableUsecase {
	return &getTimetableUsecase{
		timetableRepository: timetableRepository,
	}
}

func (u *getTimetableUsecase) GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error) {
	timetable, err := u.timetableRepository.GetByParentRouteIdAndBusStopId(ctx, parentRouteId, busStopId)
	if err != nil {
		return entity.Timetable{}, fmt.Errorf("failed to get timetable for parent route %s and bus stop %s: %w", parentRouteId, busStopId, err)
	}
	return timetable, nil
}
