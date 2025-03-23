package entity

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type TimetableEntry struct {
	DepartureTime   vo.DepartureTime
	OperationDays   vo.OperationDays
	DestinationName vo.DestinationName
}

func NewTimetableEntry(departureTime vo.DepartureTime, operationDays vo.OperationDays, destinationName vo.DestinationName) *TimetableEntry {
	return &TimetableEntry{
		DepartureTime:   departureTime,
		OperationDays:   operationDays,
		DestinationName: destinationName,
	}
}
