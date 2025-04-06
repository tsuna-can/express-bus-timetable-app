package usecase 

import (
	"context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/application/input"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
	"log"
)

type busStopsUsecase struct {
	busStopsRepository repository.BusStopsRepository
}

func NewBusStopUsecase(busStopsRepository repository.BusStopsRepository) input.BusStopsInputPort {
	return &busStopsUsecase{
		busStopsRepository: busStopsRepository,
	}
}

func (u *busStopsUsecase) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, error) {
	busStops, err := u.busStopsRepository.GetByParentRouteId(ctx, parentRouteId)
	if err != nil {
		log.Printf("Error getting bus stops: %v", err)
		return nil, err
	}
	return busStops, nil
}

