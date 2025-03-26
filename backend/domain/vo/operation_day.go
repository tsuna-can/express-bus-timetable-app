package vo

import (
	"time"
)

type OperationDay struct {
	days time.Weekday
}

func NewOperationDay(day time.Weekday) *OperationDay {
  return &OperationDay{days: day}
}

func (od *OperationDay) Value() time.Weekday {
  return od.days
}

// IntValue は曜日をIsoWeekDayのint値で返す
func (od *OperationDay) IntValue() int {
  return int(od.days)
}

