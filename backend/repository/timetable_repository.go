package repository

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase/gateway"
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

func NewTimetableRepository(db *sqlx.DB) gateway.TimetableGateway {
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
	var entries []entity.TimetableEntry

	for rows.Next() {
		var parentRouteId, parentRouteName, routeId, routeName, stopId, stopName string
		var monday, tuesday, wednesday, thursday, friday, saturday, sunday bool
		var departureTime time.Time

		if err := rows.Scan(
			&parentRouteId,
			&parentRouteName,
			&routeId,
			&routeName,
			&stopId,
			&stopName,
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

		// create departure time vo
		departureTimeVo, err := vo.NewDepartureTime(departureTime.Format("15:04"))
		if err != nil {
			log.Printf("Error creating DepartureTime: %v", err)
			return entity.Timetable{}, err
		}

		// create distination name vo
		destinationNameVo, err := vo.NewDestinationName(routeName)
		if err != nil {
			log.Printf("Error creating RouteName: %v", err)
			return entity.Timetable{}, err
		}

		// 曜日情報をtime.Weekdayのスライスに変換
		activeWeekdays := getActiveWeekdays(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

		// 曜日情報をvoに変換
		operationDays, err := vo.NewOperationDays(activeWeekdays)

		// Create timetable entry
		entry := entity.TimetableEntry{
			DepartureTime:   *departureTimeVo,
			DestinationName: *destinationNameVo,
			OperationDays:   *operationDays,
		}
		entries = append(entries, entry)
	}

	// Create timetable
	timetable := entity.Timetable{
		ParentRouteId:    parentRouteId,
		BusStopId:        busStopId,
		TimetableEntries: entries,
	}

	return timetable, nil
}

// 曜日フラグをtime.Weekdayのスライスに変換する関数
func getActiveWeekdays(monday, tuesday, wednesday, thursday, friday, saturday, sunday bool) []time.Weekday {
	weekdays := []struct {
		flag    bool
		weekday time.Weekday
	}{
		{monday, time.Monday},
		{tuesday, time.Tuesday},
		{wednesday, time.Wednesday},
		{thursday, time.Thursday},
		{friday, time.Friday},
		{saturday, time.Saturday},
		{sunday, time.Sunday},
	}

	var activeDays []time.Weekday
	for _, w := range weekdays {
		if w.flag {
			activeDays = append(activeDays, w.weekday)
		}
	}

	return activeDays
}
