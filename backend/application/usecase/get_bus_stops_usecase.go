package usecase

import (
	"context"
	"log"

	"github.com/tsuna-can/express-bus-time-table-app/backend/application/input"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

type busStopsUsecase struct {
	busStopsRepository     repository.BusStopsRepository
	parentRoutesRepository repository.ParentRoutesRepository
}

func NewBusStopUsecase(busStopsRepository repository.BusStopsRepository, parentRoutesRepository repository.ParentRoutesRepository) input.BusStopsInputPort {
	return &busStopsUsecase{
		busStopsRepository:     busStopsRepository,
		parentRoutesRepository: parentRoutesRepository,
	}
}

func (u *busStopsUsecase) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, entity.ParentRoute, error) {
	busStops, err := u.busStopsRepository.GetByParentRouteId(ctx, parentRouteId)
	if err != nil {
		log.Printf("Error getting bus stops: %v", err)
		return nil, entity.ParentRoute{}, err
	}

	parentRoute, err := u.parentRoutesRepository.GetByParentRouteId(ctx, parentRouteId)
	if err != nil {
		log.Printf("Error getting parent route: %v", err)
		return nil, entity.ParentRoute{}, err
	}

	return busStops, parentRoute, nil
}
