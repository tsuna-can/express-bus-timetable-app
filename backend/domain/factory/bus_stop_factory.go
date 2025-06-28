package factory

import (
	"fmt"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

// BusStopRawData represents raw data from database for bus stop
type BusStopRawData struct {
	BusStopId   string
	BusStopName string
}

// BusStopFactory creates BusStop entity from raw data
type BusStopFactory interface {
	ReconstructFromRawData(data BusStopRawData) (*entity.BusStop, error)
	ReconstructManyFromRawData(data []BusStopRawData) ([]entity.BusStop, error)
}

type busStopFactory struct{}

// NewBusStopFactory creates a new instance of BusStopFactory
func NewBusStopFactory() BusStopFactory {
	return &busStopFactory{}
}

// ReconstructFromRawData reconstructs a BusStop entity from raw data
func (f *busStopFactory) ReconstructFromRawData(data BusStopRawData) (*entity.BusStop, error) {
	name, err := vo.NewBusStopName(data.BusStopName)
	if err != nil {
		return nil, fmt.Errorf("failed to create BusStopName from raw data: %w", err)
	}

	return entity.NewBusStop(data.BusStopId, *name), nil
}

// ReconstructManyFromRawData reconstructs multiple BusStop entities from raw data slice
func (f *busStopFactory) ReconstructManyFromRawData(data []BusStopRawData) ([]entity.BusStop, error) {
	busStops := make([]entity.BusStop, 0, len(data))

	for i, rawData := range data {
		busStop, err := f.ReconstructFromRawData(rawData)
		if err != nil {
			return nil, fmt.Errorf("failed to reconstruct BusStop at index %d: %w", i, err)
		}
		busStops = append(busStops, *busStop)
	}

	return busStops, nil
}
