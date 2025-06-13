package response

import (
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type ParentRouteResponse struct {
  ParentRoutes []ParentRoute `json:"parent_routes"`
}

type ParentRoute struct {
  ParentRouteId string `json:"parent_route_id"`
  ParentRouteName string `json:"parent_route_name"`
}

func NewParentRoutesResponse(parentRoute []entity.ParentRoute) *ParentRouteResponse {
  parentRoutes := make([]ParentRoute, 0, len(parentRoute))
  for _, pr := range parentRoute {
    parentRoutes = append(parentRoutes, ParentRoute{
      ParentRouteId: pr.ParentRouteId,
      ParentRouteName: pr.ParentRouteName.Value(),
    })
  }
  return &ParentRouteResponse{
    ParentRoutes: parentRoutes,
  }
}

