package infrastructure

import (
	"log"

  "github.com/tsuna-can/express-bus-time-table-app/backend/handler"

	"github.com/labstack/echo/v4"
)

type Server struct {
	parentRouteHandler *handler.ParentRoutesHandler
}

func NewServer(
  parentRouteHandler *handler.ParentRoutesHandler,
) *Server {
	return &Server{
    parentRouteHandler: parentRouteHandler,
	}
}

func (s *Server) GetParentRoutes(ctx echo.Context) error {
  return s.parentRouteHandler.GetParentRoutes(ctx)
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


	e.Logger.Fatal(e.Start(":8080"))
}
