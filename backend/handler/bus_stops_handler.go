package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/request"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase"
	"net/http"
)

type BusStopsHandler struct {
	busStopsUsecase usecase.BusStopUsecase
}

func NewBusStopsHandler(busStopsUsecase usecase.BusStopUsecase) *BusStopsHandler {
	return &BusStopsHandler{
		busStopsUsecase: busStopsUsecase,
	}
}

func (h *BusStopsHandler) GetByParentRouteId(e echo.Context, parentRouteId string) error {
  req, err := request.NewGetBusStopsByParentRouteIdRequest(e)
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

