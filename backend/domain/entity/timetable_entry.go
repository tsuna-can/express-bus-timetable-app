package entity

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type TimetableEntry struct {
	DepartureTime   vo.DepartureTime
	OperationDays   map[vo.OperationDay]struct{}
	DestinationName vo.DestinationName
}

func NewTimetableEntry(
	departureTime vo.DepartureTime,
	operationDays map[vo.OperationDay]struct{},
	destinationName vo.DestinationName,
) *TimetableEntry {
	return &TimetableEntry{
		DepartureTime:   departureTime,
		OperationDays:   operationDays,
		DestinationName: destinationName,
	}
}

// OperationDays を int のスライスに変換するメソッド
func (t *TimetableEntry) OperationDaysAsIntSlice() []int {
	days := make([]int, 0, len(t.OperationDays))
	for opDay := range t.OperationDays {
		days = append(days, int(opDay.IntValue()))
	}

	return days
}
