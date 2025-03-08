package entity

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type TimetableEntry struct {
	DepartureTime   vo.DepartureTime
	DestinationName vo.DestinationName
	OperatingDays   vo.OperationDays
}

func NewTimetableEntry(departureTime vo.DepartureTime, destinationName vo.DestinationName, operationDays vo.OperationDays) *TimetableEntry {
	return &TimetableEntry{
		DepartureTime:   departureTime,
		DestinationName: destinationName,
		OperatingDays:   operationDays,
	}
}

