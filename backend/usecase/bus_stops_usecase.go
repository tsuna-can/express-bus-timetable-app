package usecase

import (
  "log"
  "context"
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

type BusStopUsecase interface {
  GetByParentRouteId(ctx context.Context, parentRouteId string) ([]model.BusStop, error)
}

type busStopUsecase struct {
  busStopRepository repository.BusStopRepository
}

func NewBusStopUsecase(busStopRepository repository.BusStopRepository) BusStopUsecase {
  return &busStopUsecase{
    busStopRepository: busStopRepository,
  }
}

func (u *busStopUsecase) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]model.BusStop, error) {
  busStops, err := u.busStopRepository.GetByParentRouteId(ctx, parentRouteId)
  if err != nil {
    log.Printf("Error getting bus stops: %v", err)
    return nil, err
  }
  return busStops, nil
}

