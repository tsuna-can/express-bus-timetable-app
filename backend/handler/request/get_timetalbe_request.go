package request

import (
	"github.com/labstack/echo/v4"
)

type TimetableRequest struct {
	ParentRouteId string `query:"parent-route-id" validate:"required"`
	BusStopId     string `query:"bus-stop-id" validate:"required"`
}

func NewTimetableRequest(c echo.Context) (*TimetableRequest, error) {
	var req TimetableRequest
	if err := c.Bind(&req); err != nil {
		return nil, err
	}
	if err := c.Validate(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
