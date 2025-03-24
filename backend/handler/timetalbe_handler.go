package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/request"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/input"
	"log"
	"net/http"
)

type TimetableHandler struct {
	timetableUsecase input.TimetableInputPort
}

func NewTimetableHandler(timetableUsecase input.TimetableInputPort) *TimetableHandler {
	return &TimetableHandler{
		timetableUsecase: timetableUsecase,
	}
}

func (h TimetableHandler) GetByParentRouteIdAndBusStopId(e echo.Context) error {
	req, err := request.NewTimetableRequest(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ctx := e.Request().Context()
	timetable, err := h.timetableUsecase.GetByParentRouteIdAndBusStopId(ctx, req.ParentRouteId, req.BusStopId)
	if err != nil {
		log.Printf("Error getting timetables: %v", err)
		return e.JSON(http.StatusInternalServerError, err)
	}

	res := response.NewTimetableResponse(timetable)
	return e.JSON(http.StatusOK, res)
}

