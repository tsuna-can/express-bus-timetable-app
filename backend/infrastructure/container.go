package infrastructure

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewServer)
	container.Provide(NewDb)

	// controllers
	container.Provide(handler.NewParentRoutesHandler)
  container.Provide(handler.NewBusStopsHandler)

	// usecases
	container.Provide(usecase.NewParentRouteUsecase)
	container.Provide(usecase.NewBusStopUsecase)

	// repositories
	container.Provide(repository.NewParentRouteRepository)
	container.Provide(repository.NewBusStopRepository)

	return container
}
