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
    r.route_id,
    r.route_name,
    s.stop_id,
    s.stop_name,
    st.departure_time,
    c.monday, c.tuesday, c.wednesday, c.thursday, c.friday, c.saturday, c.sunday
FROM StopTime st
JOIN Stop s ON st.stop_id = s.stop_id
JOIN Trip t ON st.trip_id = t.trip_id
JOIN Route r ON t.route_id = r.route_id
JOIN ParentRoute pr ON r.parent_route_id = pr.parent_route_id
JOIN Calendar c ON t.service_id = c.service_id
WHERE pr.parent_route_id = $1
AND s.stop_id = $2;
`

type TimetableRepository struct {
	db *sqlx.DB
}

func NewTimetableRepository(db *sqlx.DB) repository.TimetableRepository {
	return &TimetableRepository{db}
}

func (r *TimetableRepository) GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error) {
	rows, err := r.db.QueryContext(ctx, getTimetableQuery, parentRouteId, busStopId)
	if err != nil {
		log.Printf("Error querying timetables: %v", err)
		return entity.Timetable{}, err
	}
	defer rows.Close()

	// Create timetable entries
	var timetableModel model.Timetable
	timetableModel.ParentRouteId = parentRouteId
	timetableModel.BusStopId = busStopId

	for rows.Next() {
		var entry model.TimetableEntry
		var monday, tuesday, wednesday, thursday, friday, saturday, sunday bool
		var departureTime time.Time

		if err := rows.Scan(
			&timetableModel.ParentRouteId,
			&entry.DestinationName,
			&entry.DestinationName, // route_id (not used)
			&entry.DestinationName, // route_name
			&timetableModel.BusStopId,
			&entry.DestinationName, // stop_name
			&departureTime,
			&monday,
			&tuesday,
			&wednesday,
			&thursday,
			&friday,
			&saturday,
			&sunday,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return entity.Timetable{}, err
		}

		entry.DepartureTime = departureTime.Format("15:04")
		entry.OperationDays = getActiveWeekdaysStrings(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
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

// Helper function to convert weekday flags to string slice
func getActiveWeekdaysStrings(monday, tuesday, wednesday, thursday, friday, saturday, sunday bool) []string {
	weekdays := []struct {
		flag bool
		name string
	}{
		{monday, "monday"},
		{tuesday, "tuesday"},
		{wednesday, "wednesday"},
		{thursday, "thursday"},
		{friday, "friday"},
		{saturday, "saturday"},
		{sunday, "sunday"},
	}

	var activeDays []string
	for _, w := range weekdays {
		if w.flag {
			activeDays = append(activeDays, w.name)
		}
	}

	return activeDays
}
