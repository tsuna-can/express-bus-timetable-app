package repository

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/repository/model"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/gateway"
)

var getBusStopsQuery = `
WITH 
routes AS (
    SELECT route_id FROM route WHERE parent_route_id = $1
),
trips AS (
    SELECT trip_id FROM trip WHERE route_id IN (SELECT route_id FROM routes)
),
stop_times AS (
    SELECT stop_id FROM stoptime WHERE trip_id IN (SELECT trip_id FROM trips)
)
SELECT stop_id, stop_name FROM stop WHERE stop_id IN (SELECT stop_id FROM stop_times);
`

type BusStopsRepository struct {
	db *sqlx.DB
}

func NewBusStopRepository(db *sqlx.DB) gateway.BusStopsGateway {
	return &BusStopsRepository{db}
}

func (bsr *BusStopsRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, error) {
	rows, err := bsr.db.QueryContext(ctx, getBusStopsQuery, parentRouteId)
	if err != nil {
		log.Printf("Error querying bus stops: %v", err)
		return nil, err
	}
	defer rows.Close()

	busStops := make([]entity.BusStop, 0)
	for rows.Next() {
		var bsm model.BusStop
		if err := rows.Scan(&bsm.BusStopId, &bsm.BusStopName); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		bse, err := bsm.ToBusStop()
		if err != nil {
			log.Printf("Error converting to BusStop: %v", err)
			return nil, err
		}
		busStops = append(busStops, *bse)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, err
	}

	return busStops, nil
}

