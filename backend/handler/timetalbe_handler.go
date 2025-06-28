package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/request"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase"
)

type TimetableHandler struct {
	getTimetableUsecase usecase.GetTimetableUsecase
}

func NewTimetableHandler(getTimetableUsecase usecase.GetTimetableUsecase) *TimetableHandler {
	return &TimetableHandler{
		getTimetableUsecase: getTimetableUsecase,
	}
}

// GetByParentRouteIdAndBusStopId godoc
// @Summary Get timetable by parent route ID and bus stop ID
// @Description Get timetable by parent route ID and bus stop ID
// @Tags timetable
// @Accept json
// @Produce json
// @Param parent_route_id query string true "Parent Route ID"
// @Param bus_stop_id query string true "Bus Stop ID"
// @Success 200 {object} response.TimetableResponse
// @Router /timetable [get]
func (h TimetableHandler) GetByParentRouteIdAndBusStopId(e echo.Context) error {
	req, err := request.NewTimetableRequest(e)
	if err != nil {
		log.Printf("Error creating timetable request: %v", err)
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request parameters"})
	}

	ctx := e.Request().Context()
	timetable, err := h.getTimetableUsecase.GetByParentRouteIdAndBusStopId(ctx, req.ParentRouteId, req.BusStopId)
	if err != nil {
		log.Printf("Error getting timetables: %v", err)
		return e.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "Internal server error"})
	}

	res := response.NewTimetableResponse(timetable)
	return e.JSON(http.StatusOK, res)
}
