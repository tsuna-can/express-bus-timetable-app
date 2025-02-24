package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/request"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/input_port"
	"net/http"
)

type BusStopsHandler struct {
	busStopsUsecase input_port.BusStopsInputPort
}

func NewBusStopsHandler(busStopsUsecase input_port.BusStopsInputPort) *BusStopsHandler {
	return &BusStopsHandler{
		busStopsUsecase: busStopsUsecase,
	}
}

func (h *BusStopsHandler) GetByParentRouteId(e echo.Context) error {
  req, err := request.NewBusStopsRequest(e)
  if err != nil {
    return e.JSON(http.StatusBadRequest, err)
  }

	ctx := e.Request().Context()
	busStops, err := h.busStopsUsecase.GetByParentRouteId(ctx, req.ParentRouteId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	res:= response.NewBusStopsResponse(busStops)
	return e.JSON(http.StatusOK, res)
}

