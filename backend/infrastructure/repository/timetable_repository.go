package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/factory"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
)

const getTimetableQuery = `
SELECT
    pr.parent_route_id,
    pr.parent_route_name,
    r.route_name,
    s.stop_id,
    s.stop_name,
    st.departure_time,
    c.monday, c.tuesday, c.wednesday, c.thursday, c.friday, c.saturday, c.sunday
FROM stop_time st
JOIN stop s ON st.stop_id = s.stop_id
JOIN trip t ON st.trip_id = t.trip_id
JOIN route r ON t.route_id = r.route_id
JOIN parent_route pr ON r.parent_route_id = pr.parent_route_id
JOIN calendar c ON t.service_id = c.service_id
WHERE pr.parent_route_id = $1
AND s.stop_id = $2;
`

type TimetableRepository struct {
	db      *sqlx.DB
	factory factory.TimetableFactory
}

func NewTimetableRepository(db *sqlx.DB) repository.TimetableRepository {
	return &TimetableRepository{
		db:      db,
		factory: factory.NewTimetableFactory(),
	}
}

// FIXME : Refactor SQL and logic
func (r *TimetableRepository) GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error) {
	rows, err := r.db.QueryContext(ctx, getTimetableQuery, parentRouteId, busStopId)
	if err != nil {
		return entity.Timetable{}, fmt.Errorf("failed to query timetables: %w", err)
	}
	defer rows.Close()

	var rawData factory.TimetableRawData
	rawData.ParentRouteId = parentRouteId
	rawData.BusStopId = busStopId

	for rows.Next() {
		var entryRawData factory.TimetableEntryRawData
		var departureTime time.Time

		if err := rows.Scan(
			&rawData.ParentRouteId,
			&rawData.ParentRouteName,
			&entryRawData.DestinationName, // route_name
			&rawData.BusStopId,
			&rawData.BusStopName,
			&departureTime,
			&entryRawData.Monday,
			&entryRawData.Tuesday,
			&entryRawData.Wednesday,
			&entryRawData.Thursday,
			&entryRawData.Friday,
			&entryRawData.Saturday,
			&entryRawData.Sunday,
		); err != nil {
			return entity.Timetable{}, fmt.Errorf("failed to scan timetable row: %w", err)
		}

		entryRawData.DepartureTime = departureTime.Format("15:04")
		rawData.Entries = append(rawData.Entries, entryRawData)
	}

	if err := rows.Err(); err != nil {
		return entity.Timetable{}, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	// Convert raw data to entity using factory
	timetableEntity, err := r.factory.ReconstructFromRawData(rawData)
	if err != nil {
		return entity.Timetable{}, fmt.Errorf("failed to reconstruct timetable from raw data: %w", err)
	}

	return *timetableEntity, nil
}
