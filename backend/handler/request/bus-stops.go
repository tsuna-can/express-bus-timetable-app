package request

import (
  "github.com/labstack/echo/v4"
)

type GetBusStopsByParentRouteIdRequest struct {
  ParentRouteId string
}

func NewGetBusStopsByParentRouteIdRequest(e echo.Context) (*GetBusStopsByParentRouteIdRequest, error) {
  parentRouteId := e.Param("parentRouteId")
  return &GetBusStopsByParentRouteIdRequest{
    ParentRouteId: parentRouteId,
  }, nil
}

