package handler

import (
  "log"
	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/application/input"
  "github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
	"net/http"
)

type ParentRoutesHandler struct {
	parentRoutesUsecase input.ParentRoutesInputPort
}

func NewParentRoutesHandler(parentRoutesUsecase input.ParentRoutesInputPort) *ParentRoutesHandler {
	return &ParentRoutesHandler{
		parentRoutesUsecase: parentRoutesUsecase,
	}
}

func (h ParentRoutesHandler) GetParentRoutes(e echo.Context) error {
	ctx := e.Request().Context()
	parentRoutes, err := h.parentRoutesUsecase.GetAll(ctx)
	if err != nil {
    log.Printf("Error getting parent routes: %v", err)
		return e.JSON(http.StatusInternalServerError, err)
	}

  response := response.NewParentRoutesResponse(parentRoutes)
  return e.JSON(http.StatusOK, response)
}

