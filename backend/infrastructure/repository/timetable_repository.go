package repository

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/repository"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/model"
)

var getTimetableQuery = `
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
		log.Printf("Error querying timetables: %v", err)
		return entity.Timetable{}, err
	}
	defer rows.Close()

	// Create timetable entries
	var timetableModel model.Timetable

	for rows.Next() {
		var entry model.TimetableEntry
		// var monday, tuesday, wednesday, thursday, friday, saturday, sunday bool
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
			log.Printf("Error scanning row: %v", err)
			return entity.Timetable{}, err
		}

		entry.DepartureTime = departureTime.Format("15:04")
		timetableModel.TimetableEntry = append(timetableModel.TimetableEntry, entry)
	}

	// Convert model to entity
	timetableEntity, err := timetableModel.ToTimetable()
	if err != nil {
		log.Printf("Error converting model to entity: %v", err)
		return entity.Timetable{}, err
	}

	return *timetableEntity, nil
}
