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

	// handlers
	container.Provide(handler.NewParentRoutesHandler)
	container.Provide(handler.NewBusStopsHandler)
	container.Provide(handler.NewTimetableHandler)

	// usecases
	container.Provide(usecase.NewGetParentRoutesUsecase)
	container.Provide(usecase.NewGetBusStopsUsecase)
	container.Provide(usecase.NewGetTimetableUsecase)

	// repositories
	container.Provide(repository.NewParentRoutesRepository)
	container.Provide(repository.NewBusStopsRepository)
	container.Provide(repository.NewTimetableRepository)

	return container
}
