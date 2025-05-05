package model

import (
	"fmt"
	"time"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type Timetable struct {
	ParentRouteId   string
	ParentRouteName string
	BusStopName     string
	BusStopId       string
	TimetableEntry  []TimetableEntry
}

type TimetableEntry struct {
	DepartureTime   string
	DestinationName string
	Monday          bool
	Tuesday         bool
	Wednesday       bool
	Thursday        bool
	Friday          bool
	Saturday        bool
	Sunday          bool
}

// ToTimetable converts Timetable to entity.Timetable
func (t *Timetable) ToTimetable() (*entity.Timetable, error) {
	timetableEntries := make([]entity.TimetableEntry, len(t.TimetableEntry))

	parentRouteName, err := vo.NewParentRouteName(t.ParentRouteName)
	if err != nil {
		return nil, fmt.Errorf("failed to create ParentRouteName: %w", err)
	}
	busStopName, err := vo.NewBusStopName(t.BusStopName)
	if err != nil {
		return nil, fmt.Errorf("failed to create BusStopName: %w", err)
	}

	for i, te := range t.TimetableEntry {
		departureTime, err := vo.NewDepartureTime(te.DepartureTime)
		if err != nil {
			return nil, fmt.Errorf("failed to create DepartureTime: %w", err)
		}

		destinationName, err := vo.NewDestinationName(te.DestinationName)
		if err != nil {
			return nil, fmt.Errorf("failed to create DestinationName: %w", err)
		}

		operationDays := make(map[vo.OperationDay]struct{})
		addOperationDay := func(day time.Weekday) {
			opDay := vo.NewOperationDay(day) // *vo.OperationDay（ポインタ）
			if opDay != nil {                // nil チェック
				operationDays[*opDay] = struct{}{} // ポインタをデリファレンスして格納
			}
		}

		if te.Monday {
			addOperationDay(time.Monday)
		}
		if te.Tuesday {
			addOperationDay(time.Tuesday)
		}
		if te.Wednesday {
			addOperationDay(time.Wednesday)
		}
		if te.Thursday {
			addOperationDay(time.Thursday)
		}
		if te.Friday {
			addOperationDay(time.Friday)
		}
		if te.Saturday {
			addOperationDay(time.Saturday)
		}
		if te.Sunday {
			addOperationDay(time.Sunday)
		}

		timetableEntries[i] = entity.TimetableEntry{
			DepartureTime:   *departureTime,
			DestinationName: *destinationName,
			OperationDays:   operationDays,
		}
	}

	return entity.NewTimetable(
		t.ParentRouteId, // t.ParentRouteIdを使用
		*parentRouteName,
		t.BusStopId, // t.BusStopIdを使用
		*busStopName,
		timetableEntries,
	), nil
}
