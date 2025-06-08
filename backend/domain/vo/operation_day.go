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

// IntValue は曜日をISOフォーマットのint値で返す（月曜日=1、...日曜日=7）
func (od *OperationDay) IntValue() int {
	// Go標準では日曜日が0なので、ISO表記（日曜日=7）に変換する
	if od.days == time.Sunday {
		return 7
	}
	return int(od.days)
}
