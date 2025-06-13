package infrastructure

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/application/usecase"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewServer)
	container.Provide(NewDb)

	// handlers
	container.Provide(handler.NewParentRoutesHandler)
	container.Provide(handler.NewBusStopsHandler)
	container.Provide(handler.NewTimetableHandler)

	// usecases
	container.Provide(usecase.NewParentRoutesUsecase)
	container.Provide(usecase.NewBusStopUsecase)
	container.Provide(usecase.NewTimetableUsecase)

	// repositories
	container.Provide(repository.NewParentRoutesRepository)
	container.Provide(repository.NewBusStopRepository)
	container.Provide(repository.NewTimetableRepository)

	return container
}
