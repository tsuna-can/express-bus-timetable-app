package vo

import (
	"testing"
)

func TestNewBusStopName(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "valid bus stop name",
			input:       "東京駅前",
			expectError: false,
		},
		{
			name:        "empty string",
			input:       "",
			expectError: true,
		},
		{
			name:        "only whitespace",
			input:       "   ",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBusStopName(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("NewBusStopName() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && got.Value() != tt.input {
				t.Errorf("BusStopName.Value() = %v, want %v", got.Value(), tt.input)
			}
		})
	}
}
