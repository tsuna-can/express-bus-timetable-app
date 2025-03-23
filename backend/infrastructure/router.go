package infrastructure

import (
	"log"

	"github.com/tsuna-can/express-bus-time-table-app/backend/handler"

	"github.com/labstack/echo/v4"
)

type Server struct {
	parentRouteHandler *handler.ParentRoutesHandler
	busStopHandler     *handler.BusStopsHandler
  timeTableHandler   *handler.TimetableHandler
}

func NewServer(
	parentRouteHandler *handler.ParentRoutesHandler,
	busStopHandler *handler.BusStopsHandler,
  timeTableHandler *handler.TimetableHandler,
) *Server {
	return &Server{
		parentRouteHandler: parentRouteHandler,
		busStopHandler:     busStopHandler,
    timeTableHandler:   timeTableHandler,
	}
}

func (s *Server) GetParentRoutes(ctx echo.Context) error {
	return s.parentRouteHandler.GetParentRoutes(ctx)
}

func (s *Server) GetBusStopsByParentRouteId(ctx echo.Context) error {
	return s.busStopHandler.GetByParentRouteId(ctx)
}

func (s *Server) GetTimetableByParentRouteIdAndBusStopId(ctx echo.Context) error {
	return s.timeTableHandler.GetByParentRouteIdAndBusStopId(ctx)
}

func InitRouter() {
	e := echo.New()

	container := BuildContainer()

	var server *Server
	if err := container.Invoke(func(s *Server) {
		server = s
	}); err != nil {
		log.Fatalf("Error resolving dependencies: %v", err)
	}

	e.GET("/parent-routes", server.GetParentRoutes)
	e.GET("/bus-stops", server.GetBusStopsByParentRouteId)
	e.GET("/timetable", server.GetTimetableByParentRouteIdAndBusStopId)

	e.Logger.Fatal(e.Start(":8080"))
}
