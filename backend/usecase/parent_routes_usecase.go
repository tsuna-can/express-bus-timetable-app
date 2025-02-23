package usecase

import (
  "log"
  "context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/model"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

type ParentRouteUsecase interface {
	GetAll(ctx context.Context) ([]model.ParentRoute, error)
}

type parentRouteUsecase struct {
	parentRouteRepository repository.ParentRouteRepository
}

func NewParentRouteUsecase(parentRouteRepository repository.ParentRouteRepository) ParentRouteUsecase {
	return &parentRouteUsecase{
		parentRouteRepository: parentRouteRepository,
	}
}

func (u *parentRouteUsecase) GetAll(ctx context.Context) ([]model.ParentRoute, error) {
  parentRoutes, err := u.parentRouteRepository.GetAll(ctx)
  if err != nil {
    log.Printf("Error getting parent routes: %v", err)
    return nil, err
  }
  return parentRoutes, nil
}
