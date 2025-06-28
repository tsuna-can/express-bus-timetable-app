package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/request"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase"
)

type BusStopsHandler struct {
	getBusStopsUsecase usecase.GetBusStopsUsecase
}

func NewBusStopsHandler(getBusStopsUsecase usecase.GetBusStopsUsecase) *BusStopsHandler {
	return &BusStopsHandler{
		getBusStopsUsecase: getBusStopsUsecase,
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
		log.Printf("Error creating bus stops request: %v", err)
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request parameters"})
	}

	ctx := e.Request().Context()
	busStops, parentRoute, err := h.getBusStopsUsecase.GetByParentRouteId(ctx, req.ParentRouteId)
	if err != nil {
		log.Printf("Error getting bus stops: %v", err)
		return e.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "Internal server error"})
	}

	res := response.NewBusStopsResponse(busStops, parentRoute)
	return e.JSON(http.StatusOK, res)
}
