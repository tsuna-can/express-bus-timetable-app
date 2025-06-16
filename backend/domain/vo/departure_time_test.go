package vo

import (
	"testing"
)

func TestNewDepartureTime(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "valid time format",
			input:       "12:30",
			expectError: false,
		},
		{
			name:        "valid time format - early morning",
			input:       "05:15",
			expectError: false,
		},
		{
			name:        "valid time format - late night",
			input:       "23:59",
			expectError: false,
		},
		{
			name:        "invalid format - no colon",
			input:       "1230",
			expectError: true,
		},
		{
			name:        "invalid format - too many parts",
			input:       "12:30:00",
			expectError: true,
		},
		{
			name:        "invalid format - letters",
			input:       "ab:cd",
			expectError: true,
		},
		{
			name:        "invalid hours - too large",
			input:       "24:00",
			expectError: true,
		},
		{
			name:        "invalid hours - negative",
			input:       "-1:30",
			expectError: true,
		},
		{
			name:        "invalid minutes - too large",
			input:       "12:60",
			expectError: true,
		},
		{
			name:        "invalid minutes - negative",
			input:       "12:-1",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDepartureTime(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("NewDepartureTime() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && got.Value() != tt.input {
				t.Errorf("DepartureTime.Value() = %v, want %v", got.Value(), tt.input)
			}
		})
	}
}

func TestIsValidTimeFormat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid format",
			input: "12:30",
			want:  true,
		},
		{
			name:  "invalid - no colon",
			input: "1230",
			want:  false,
		},
		{
			name:  "invalid - multiple colons",
			input: "12:30:45",
			want:  false,
		},
		{
			name:  "invalid - non-numeric",
			input: "ab:cd",
			want:  false,
		},
		{
			name:  "invalid - hours too large",
			input: "24:00",
			want:  false,
		},
		{
			name:  "invalid - minutes too large",
			input: "12:60",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidTimeFormat(tt.input); got != tt.want {
				t.Errorf("isValidTimeFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
