package interactor

import (
	"context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/gateway"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/input"
	"log"
)

type parentRoutesUsecase struct {
	parentRoutesRepository gateway.ParentRoutesGateway
}

func NewParentRoutesUsecase(parentRoutesRepository gateway.ParentRoutesGateway) input.ParentRoutesInputPort {
	return &parentRoutesUsecase{
		parentRoutesRepository: parentRoutesRepository,
	}
}

func (u *parentRoutesUsecase) GetAll(ctx context.Context) ([]entity.ParentRoute, error) {
	parentRoutes, err := u.parentRoutesRepository.GetAll(ctx)
	if err != nil {
		log.Printf("Error getting parent routes: %v", err)
		return nil, err
	}
	return parentRoutes, nil
}
