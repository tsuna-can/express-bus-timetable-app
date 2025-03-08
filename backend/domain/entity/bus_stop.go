package entity

import (
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)


type BusStop struct {
	BusStopId   string
	BusStopName vo.BusStopName
}

func NewBusStop(busStopId string, busStopName vo.BusStopName) *BusStop {
	return &BusStop{
		BusStopId:   busStopId,
		BusStopName: busStopName,
	}
}

