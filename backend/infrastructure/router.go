package infrastructure

import (
	"log"

  "github.com/tsuna-can/express-bus-time-table-app/backend/handler"

	"github.com/labstack/echo/v4"
)

type Server struct {
	parentRouteHandler *handler.ParentRoutesHandler
  busStopHandler *handler.BusStopsHandler
}

func NewServer(
  parentRouteHandler *handler.ParentRoutesHandler,
  busStopHandler *handler.BusStopsHandler,
) *Server {
	return &Server{
    parentRouteHandler: parentRouteHandler,
    busStopHandler: busStopHandler,
	}
}

func (s *Server) GetParentRoutes(ctx echo.Context) error {
  return s.parentRouteHandler.GetParentRoutes(ctx)
}

func (s *Server) GetBusStopsByParentRouteId(ctx echo.Context) error {
	parentRouteId := ctx.QueryParam("parent_route_id") // クエリパラメータから取得
	return s.busStopHandler.GetByParentRouteId(ctx, parentRouteId)
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


	e.Logger.Fatal(e.Start(":8080"))
}
