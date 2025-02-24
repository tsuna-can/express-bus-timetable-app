package entity

type BusStop struct {
	BusStopId   string
	BusStopName string
}

func NewBusStop(busStopId, busStopName string) *BusStop {
	return &BusStop{
		BusStopId:   busStopId,
		BusStopName: busStopName,
	}
}

