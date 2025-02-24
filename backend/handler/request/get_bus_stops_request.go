package request

import (
  "github.com/labstack/echo/v4"
)

type BusStopsRequest struct {
  ParentRouteId string
}

func NewBusStopsRequest(e echo.Context) (*BusStopsRequest, error) {
  parentRouteId := e.QueryParam("parent-route-id")
  return &BusStopsRequest{
    ParentRouteId: parentRouteId,
  }, nil
}

