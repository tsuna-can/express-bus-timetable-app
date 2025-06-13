package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/application/input"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/request"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
)

type BusStopsHandler struct {
	busStopsUsecase input.BusStopsInputPort
}

func NewBusStopsHandler(busStopsUsecase input.BusStopsInputPort) *BusStopsHandler {
	return &BusStopsHandler{
		busStopsUsecase: busStopsUsecase,
	}
}

// GetByParentRouteId godoc
// @Summary Get bus stops by parent route ID
// @Description Get bus stops by parent route ID
// @Tags bus_stops
// @Accept json
// @Produce json
// @Param parent_route_id query string true "Parent Route ID"
// @Success 200 {array} response.BusStopResponse
// @Router /bus-stops [get]
func (h *BusStopsHandler) GetByParentRouteId(e echo.Context) error {
	req, err := request.NewBusStopsRequest(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ctx := e.Request().Context()
	busStops, parentRoute, err := h.busStopsUsecase.GetByParentRouteId(ctx, req.ParentRouteId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	res := response.NewBusStopsResponse(busStops, parentRoute)
	return e.JSON(http.StatusOK, res)
}
