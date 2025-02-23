package model

import (
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type BusStop struct {
  bus_stop_id string
  bus_stop_name string
}

func (bs *BusStop) ToBusStop() *model.BusStop {
  return model.NewBusStop(
    bs.bus_stop_id,
    bs.bus_stop_name,
  )
}

