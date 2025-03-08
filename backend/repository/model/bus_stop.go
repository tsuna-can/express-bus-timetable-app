package model

import (
	"fmt"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type BusStop struct {
	BusStopId   string
	BusStopName string
}

func (bs *BusStop) ToBusStop() (*entity.BusStop, error) {
	name, err := vo.NewBusStopName(bs.BusStopName)
	if err != nil {
		return nil, fmt.Errorf("failed to create BusStopName: %w", err)
	}

	return entity.NewBusStop(
		bs.BusStopId,
		*name,
	), nil
}
