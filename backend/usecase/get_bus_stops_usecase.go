package usecase

import (
	"context"
	"fmt"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

type GetBusStopsUsecase interface {
	GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, entity.ParentRoute, error)
}

type getBusStopsUsecase struct {
	busStopsRepository     repository.BusStopsRepository
	parentRoutesRepository repository.ParentRoutesRepository
}

func NewGetBusStopsUsecase(busStopsRepository repository.BusStopsRepository, parentRoutesRepository repository.ParentRoutesRepository) GetBusStopsUsecase {
	return &getBusStopsUsecase{
		busStopsRepository:     busStopsRepository,
		parentRoutesRepository: parentRoutesRepository,
	}
}

func (u *getBusStopsUsecase) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, entity.ParentRoute, error) {
	busStops, err := u.busStopsRepository.GetByParentRouteId(ctx, parentRouteId)
	if err != nil {
		return nil, entity.ParentRoute{}, fmt.Errorf("failed to get bus stops for parent route %s: %w", parentRouteId, err)
	}

	parentRoute, err := u.parentRoutesRepository.GetByParentRouteId(ctx, parentRouteId)
	if err != nil {
		return nil, entity.ParentRoute{}, fmt.Errorf("failed to get parent route %s: %w", parentRouteId, err)
	}

	return busStops, parentRoute, nil
}
