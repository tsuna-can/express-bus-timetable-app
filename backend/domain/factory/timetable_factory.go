package factory

import (
	"fmt"
	"time"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

// TimetableRawData represents raw data from database for timetable
type TimetableRawData struct {
	ParentRouteId   string
	ParentRouteName string
	BusStopId       string
	BusStopName     string
	Entries         []TimetableEntryRawData
}

// TimetableEntryRawData represents raw data for a single timetable entry
type TimetableEntryRawData struct {
	DepartureTime   string // Already formatted as "15:04"
	DestinationName string
	// Operation days as individual bool fields (simplifies SQL scanning)
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool
	Sunday    bool
}

// TimetableFactory creates Timetable entity from raw data
type TimetableFactory interface {
	ReconstructFromRawData(data TimetableRawData) (*entity.Timetable, error)
}

type timetableFactory struct{}

// NewTimetableFactory creates a new instance of TimetableFactory
func NewTimetableFactory() TimetableFactory {
	return &timetableFactory{}
}

// ReconstructFromRawData reconstructs a Timetable entity from raw data
func (f *timetableFactory) ReconstructFromRawData(data TimetableRawData) (*entity.Timetable, error) {
	parentRouteName, err := vo.NewParentRouteName(data.ParentRouteName)
	if err != nil {
		return nil, fmt.Errorf("failed to create ParentRouteName from raw data: %w", err)
	}

	busStopName, err := vo.NewBusStopName(data.BusStopName)
	if err != nil {
		return nil, fmt.Errorf("failed to create BusStopName from raw data: %w", err)
	}

	timetableEntries := make([]entity.TimetableEntry, len(data.Entries))
	for i, entryData := range data.Entries {
		entry, err := f.createTimetableEntry(entryData)
		if err != nil {
			return nil, fmt.Errorf("failed to create TimetableEntry at index %d: %w", i, err)
		}
		timetableEntries[i] = *entry
	}

	return entity.NewTimetable(
		data.ParentRouteId,
		*parentRouteName,
		data.BusStopId,
		*busStopName,
		timetableEntries,
	), nil
}

// createTimetableEntry creates a single TimetableEntry from raw data
func (f *timetableFactory) createTimetableEntry(data TimetableEntryRawData) (*entity.TimetableEntry, error) {
	departureTime, err := vo.NewDepartureTime(data.DepartureTime)
	if err != nil {
		return nil, fmt.Errorf("failed to create DepartureTime from raw data: %w", err)
	}

	destinationName, err := vo.NewDestinationName(data.DestinationName)
	if err != nil {
		return nil, fmt.Errorf("failed to create DestinationName from raw data: %w", err)
	}

	operationDays := f.createOperationDays(data)

	return &entity.TimetableEntry{
		DepartureTime:   *departureTime,
		DestinationName: *destinationName,
		OperationDays:   operationDays,
	}, nil
}

// createOperationDays converts boolean fields to OperationDays map
func (f *timetableFactory) createOperationDays(data TimetableEntryRawData) map[vo.OperationDay]struct{} {
	operationDays := make(map[vo.OperationDay]struct{})

	dayMapping := []struct {
		enabled bool
		weekday time.Weekday
	}{
		{data.Monday, time.Monday},
		{data.Tuesday, time.Tuesday},
		{data.Wednesday, time.Wednesday},
		{data.Thursday, time.Thursday},
		{data.Friday, time.Friday},
		{data.Saturday, time.Saturday},
		{data.Sunday, time.Sunday},
	}

	for _, mapping := range dayMapping {
		if mapping.enabled {
			if opDay := vo.NewOperationDay(mapping.weekday); opDay != nil {
				operationDays[*opDay] = struct{}{}
			}
		}
	}

	return operationDays
}
