package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/factory"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

// FIXME : Compare below query with the one that uses JOIN statements
const getBusStopsQuery = `
WITH 
routes AS (
    SELECT route_id FROM route WHERE parent_route_id = $1
),
trips AS (
    SELECT trip_id FROM trip WHERE route_id IN (SELECT route_id FROM routes)
),
stop_times AS (
    SELECT stop_id FROM stop_time WHERE trip_id IN (SELECT trip_id FROM trips)
)
SELECT stop_id, stop_name FROM stop WHERE stop_id IN (SELECT stop_id FROM stop_times);
`

type BusStopsRepository struct {
	db      *sqlx.DB
	factory factory.BusStopFactory
}

func NewBusStopsRepository(db *sqlx.DB) repository.BusStopsRepository {
	return &BusStopsRepository{
		db:      db,
		factory: factory.NewBusStopFactory(),
	}
}

func (bsr *BusStopsRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, error) {
	rows, err := bsr.db.QueryContext(ctx, getBusStopsQuery, parentRouteId)
	if err != nil {
		return nil, fmt.Errorf("failed to query bus stops: %w", err)
	}
	defer rows.Close()

	var rawDataList []factory.BusStopRawData
	for rows.Next() {
		var rawData factory.BusStopRawData
		if err := rows.Scan(&rawData.BusStopId, &rawData.BusStopName); err != nil {
			return nil, fmt.Errorf("failed to scan bus stop row: %w", err)
		}
		rawDataList = append(rawDataList, rawData)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	// Convert raw data to entities using factory
	busStops, err := bsr.factory.ReconstructManyFromRawData(rawDataList)
	if err != nil {
		return nil, fmt.Errorf("failed to reconstruct bus stops from raw data: %w", err)
	}

	return busStops, nil
}
