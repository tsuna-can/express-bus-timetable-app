package model

import (
	"fmt"
	"time"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type Timetable struct {
	ParentRouteId  string
	BusStopId      string
	TimetableEntry []TimetableEntry
}

type TimetableEntry struct {
	DepartureTime   string
	DestinationName string
	OperationDays   []string
}

func stringToWeekday(day string) (time.Weekday, error) {
	switch day {
	case "monday":
		return time.Monday, nil
	case "tuesday":
		return time.Tuesday, nil
	case "wednesday":
		return time.Wednesday, nil
	case "thursday":
		return time.Thursday, nil
	case "friday":
		return time.Friday, nil
	case "saturday":
		return time.Saturday, nil
	case "sunday":
		return time.Sunday, nil
	default:
		return time.Sunday, fmt.Errorf("invalid day: %s", day)
	}
}

// ToTimetable converts Timetable to entity.Timetable
func (t *Timetable) ToTimetable() (*entity.Timetable, error) {
	timetableEntries := make([]entity.TimetableEntry, len(t.TimetableEntry))
	for i, te := range t.TimetableEntry {
		departureTime, err := vo.NewDepartureTime(te.DepartureTime)
		if err != nil {
			return nil, fmt.Errorf("failed to create DepartureTime: %w", err)
		}

		destinationName, err := vo.NewDestinationName(te.DestinationName)
		if err != nil {
			return nil, fmt.Errorf("failed to create DestinationName: %w", err)
		}

		days := make([]time.Weekday, len(te.OperationDays))
		for i, day := range te.OperationDays {
			weekday, err := stringToWeekday(day)
			if err != nil {
				return nil, fmt.Errorf("failed to convert day to time.Weekday: %w", err)
			}
			days[i] = weekday
		}

		operationDays, err := vo.NewOperationDays(days)
		if err != nil {
			return nil, fmt.Errorf("failed to create OperationDays: %w", err)
		}

		timetableEntries[i] = entity.TimetableEntry{
			DepartureTime:   *departureTime,
			DestinationName: *destinationName,
			OperationDays:   *operationDays,
		}
	}

	return entity.NewTimetable(
		t.ParentRouteId, // t.ParentRouteIdを使用
		t.BusStopId,     // t.BusStopIdを使用
		timetableEntries,
	), nil
}
