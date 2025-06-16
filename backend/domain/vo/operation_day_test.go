package vo

import (
	"testing"
	"time"
)

func TestNewOperationDay(t *testing.T) {
	tests := []struct {
		name     string
		day      time.Weekday
		wantDay  time.Weekday
		wantInt  int
		isoValue int
	}{
		{
			name:     "Monday",
			day:      time.Monday,
			wantDay:  time.Monday,
			wantInt:  1,
			isoValue: 1,
		},
		{
			name:     "Tuesday",
			day:      time.Tuesday,
			wantDay:  time.Tuesday,
			wantInt:  2,
			isoValue: 2,
		},
		{
			name:     "Wednesday",
			day:      time.Wednesday,
			wantDay:  time.Wednesday,
			wantInt:  3,
			isoValue: 3,
		},
		{
			name:     "Thursday",
			day:      time.Thursday,
			wantDay:  time.Thursday,
			wantInt:  4,
			isoValue: 4,
		},
		{
			name:     "Friday",
			day:      time.Friday,
			wantDay:  time.Friday,
			wantInt:  5,
			isoValue: 5,
		},
		{
			name:     "Saturday",
			day:      time.Saturday,
			wantDay:  time.Saturday,
			wantInt:  6,
			isoValue: 6,
		},
		{
			name:     "Sunday",
			day:      time.Sunday,
			wantDay:  time.Sunday,
			wantInt:  0,
			isoValue: 7, // ISO format has Sunday as 7
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOperationDay(tt.day)
			if got.Value() != tt.wantDay {
				t.Errorf("OperationDay.Value() = %v, want %v", got.Value(), tt.wantDay)
			}
			if got.IntValue() != tt.isoValue {
				t.Errorf("OperationDay.IntValue() = %v, want %v", got.IntValue(), tt.isoValue)
			}
		})
	}
}
