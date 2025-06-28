package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/model"
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
	db *sqlx.DB
}

func NewBusStopsRepository(db *sqlx.DB) repository.BusStopsRepository {
	return &BusStopsRepository{db}
}

func (bsr *BusStopsRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, error) {
	rows, err := bsr.db.QueryContext(ctx, getBusStopsQuery, parentRouteId)
	if err != nil {
		return nil, fmt.Errorf("failed to query bus stops: %w", err)
	}
	defer rows.Close()

	busStops := make([]entity.BusStop, 0)
	for rows.Next() {
		var bsm model.BusStop
		if err := rows.Scan(&bsm.BusStopId, &bsm.BusStopName); err != nil {
			return nil, fmt.Errorf("failed to scan bus stop row: %w", err)
		}

		bse, err := bsm.ToBusStop()
		if err != nil {
			return nil, fmt.Errorf("failed to convert to BusStop: %w", err)
		}

		busStops = append(busStops, *bse)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	return busStops, nil
}
