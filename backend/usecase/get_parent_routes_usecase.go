package usecase

import (
	"context"
	"fmt"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

type GetParentRoutesUsecase interface {
	GetAll(ctx context.Context) ([]entity.ParentRoute, error)
}

type getParentRoutesUsecase struct {
	parentRoutesRepository repository.ParentRoutesRepository
}

func NewGetParentRoutesUsecase(parentRoutesRepository repository.ParentRoutesRepository) GetParentRoutesUsecase {
	return &getParentRoutesUsecase{
		parentRoutesRepository: parentRoutesRepository,
	}
}

func (u *getParentRoutesUsecase) GetAll(ctx context.Context) ([]entity.ParentRoute, error) {
	parentRoutes, err := u.parentRoutesRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all parent routes: %w", err)
	}
	return parentRoutes, nil
}
