package vo

import (
	"testing"
)

func TestNewParentRouteName(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "valid parent route name",
			input:       "東京駅～横浜駅",
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
			got, err := NewParentRouteName(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("NewParentRouteName() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && got.Value() != tt.input {
				t.Errorf("ParentRouteName.Value() = %v, want %v", got.Value(), tt.input)
			}
		})
	}
}
