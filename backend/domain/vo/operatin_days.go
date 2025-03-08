package vo

import (
	"fmt"
	"time"
)

type OperationDays struct {
	days map[time.Weekday]struct{}
}

func NewOperationDays(days ...time.Weekday) (OperationDays, error) {
	if len(days) == 0 {
		return OperationDays{}, fmt.Errorf("OperationDays must have at least one operation day")
	}

	uniqueDays := make(map[time.Weekday]struct{})
	for _, day := range days {
		uniqueDays[day] = struct{}{}
	}

	return OperationDays{days: uniqueDays}, nil
}

// Contains は指定した曜日が有効日かどうかを判定する
func (vd OperationDays) Contains(day time.Weekday) bool {
	_, exists := vd.days[day]
	return exists
}

// Days は有効な曜日のスライスを返す
func (vd OperationDays) Days() []time.Weekday {
	result := make([]time.Weekday, 0, len(vd.days))
	for day := range vd.days {
		result = append(result, day)
	}
	return result
}

// String は曜日をカンマ区切りの文字列として返す
func (od OperationDays) String() string {
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	result := []string{}
	for _, day := range od.Days() {
		result = append(result, weekdays[day])
	}
	return fmt.Sprintf("[%s]", fmt.Sprint(result))
}

