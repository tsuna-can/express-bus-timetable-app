package request

import (
	"github.com/labstack/echo/v4"
)

type BusStopsRequest struct {
	ParentRouteId string `query:"parent-route-id" validate:"required"`
}

func NewBusStopsRequest(c echo.Context) (*BusStopsRequest, error) {
	var req BusStopsRequest
	if err := c.Bind(&req); err != nil {
		return nil, err
	}
	if err := c.Validate(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
