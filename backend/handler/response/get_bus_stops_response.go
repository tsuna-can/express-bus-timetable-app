package response

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type BusStopResponse struct {
	ParentRouteId   string    `json:"parent_route_id"`
	ParentRouteName string    `json:"parent_route_name"`
	BusStops        []BusStop `json:"bus_stops"`
}

type BusStop struct {
	BusStopId   string `json:"bus_stop_id"`
	BusStopName string `json:"bus_stop_name"`
}

func NewBusStopsResponse(busStops []entity.BusStop, parentRoute entity.ParentRoute) *BusStopResponse {
	busStopResponses := make([]BusStop, 0, len(busStops))
	for _, bs := range busStops {
		busStopResponses = append(busStopResponses, BusStop{
			BusStopId:   bs.BusStopId,
			BusStopName: bs.BusStopName.Value(),
		})
	}
	return &BusStopResponse{
		ParentRouteId:   parentRoute.ParentRouteId,
		ParentRouteName: parentRoute.ParentRouteName.Value(),
		BusStops:        busStopResponses,
	}
}
