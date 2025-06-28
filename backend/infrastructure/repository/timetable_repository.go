package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/model"
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
	db *sqlx.DB
}

func NewTimetableRepository(db *sqlx.DB) repository.TimetableRepository {
	return &TimetableRepository{db}
}

// FIXME : Refactor SQL and logic
func (r *TimetableRepository) GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error) {
	rows, err := r.db.QueryContext(ctx, getTimetableQuery, parentRouteId, busStopId)
	if err != nil {
		return entity.Timetable{}, fmt.Errorf("failed to query timetables: %w", err)
	}
	defer rows.Close()

	// Create timetable entries
	var timetableModel model.Timetable

	for rows.Next() {
		var entry model.TimetableEntry
		var departureTime time.Time

		if err := rows.Scan(
			&timetableModel.ParentRouteId,
			&timetableModel.ParentRouteName,
			&entry.DestinationName, // route_name
			&timetableModel.BusStopId,
			&timetableModel.BusStopName,
			&departureTime,
			&entry.Monday,
			&entry.Tuesday,
			&entry.Wednesday,
			&entry.Thursday,
			&entry.Friday,
			&entry.Saturday,
			&entry.Sunday,
		); err != nil {
			return entity.Timetable{}, fmt.Errorf("failed to scan timetable row: %w", err)
		}

		entry.DepartureTime = departureTime.Format("15:04")
		timetableModel.TimetableEntry = append(timetableModel.TimetableEntry, entry)
	}

	if err := rows.Err(); err != nil {
		return entity.Timetable{}, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	// Convert model to entity
	timetableEntity, err := timetableModel.ToTimetable()
	if err != nil {
		return entity.Timetable{}, fmt.Errorf("failed to convert model to entity: %w", err)
	}

	return *timetableEntity, nil
}
