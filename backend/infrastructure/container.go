package infrastructure

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler"
	"github.com/tsuna-can/express-bus-time-table-app/backend/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/interactor"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewServer)
	container.Provide(NewDb)

	// handlers
	container.Provide(handler.NewParentRoutesHandler)
	container.Provide(handler.NewBusStopsHandler)

	// usecases
	container.Provide(interactor.NewParentRoutesUsecase)
	container.Provide(interactor.NewBusStopUsecase)

	// repositories
	container.Provide(repository.NewParentRoutesRepository)
	container.Provide(repository.NewBusStopRepository)

	return container
}

