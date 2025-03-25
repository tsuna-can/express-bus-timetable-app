package request

import (
  "github.com/labstack/echo/v4"
)

type TimetableRequest struct {
  ParentRouteId string
  BusStopId string
}

func NewTimetableRequest(e echo.Context) (*TimetableRequest, error) {
  parentRouteId := e.QueryParam("parent-route-id")
  busStopId := e.QueryParam("bus-stop-id")

  return &TimetableRequest{
    ParentRouteId: parentRouteId,
    BusStopId: busStopId,
  }, nil
}
