package vo

import (
	"testing"
)

func TestNewDestinationName(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "valid destination name",
			input:       "渋谷駅",
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
			got, err := NewDestinationName(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("NewDestinationName() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && got.Value() != tt.input {
				t.Errorf("DestinationName.Value() = %v, want %v", got.Value(), tt.input)
			}
		})
	}
}
