package infrastructure

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCustomValidator_Validate(t *testing.T) {
	validator := NewCustomValidator()

	t.Run("Valid struct should pass validation", func(t *testing.T) {
		type TestStruct struct {
			Name  string `validate:"required"`
			Email string `validate:"required,email"`
		}

		validStruct := TestStruct{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		err := validator.Validate(validStruct)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("Invalid struct should fail validation", func(t *testing.T) {
		type TestStruct struct {
			Name  string `validate:"required"`
			Email string `validate:"required,email"`
		}

		invalidStruct := TestStruct{
			Name:  "",
			Email: "invalid-email",
		}

		err := validator.Validate(invalidStruct)
		if err == nil {
			t.Error("Expected validation error, got nil")
		}

		// Check if it's an Echo HTTP error
		if httpErr, ok := err.(*echo.HTTPError); ok {
			if httpErr.Code != 400 {
				t.Errorf("Expected status code 400, got: %d", httpErr.Code)
			}
		} else {
			t.Error("Expected Echo HTTP error")
		}
	})

	t.Run("Missing required field should fail validation", func(t *testing.T) {
		type TestStruct struct {
			Name string `validate:"required"`
		}

		invalidStruct := TestStruct{
			Name: "",
		}

		err := validator.Validate(invalidStruct)
		if err == nil {
			t.Error("Expected validation error, got nil")
		}

		// Check if it's an Echo HTTP error with proper message
		if httpErr, ok := err.(*echo.HTTPError); ok {
			if httpErr.Code != 400 {
				t.Errorf("Expected status code 400, got: %d", httpErr.Code)
			}

			// Check if the message contains validation details
			if message, ok := httpErr.Message.(map[string]interface{}); ok {
				if message["message"] != "Validation failed" {
					t.Errorf("Expected message 'Validation failed', got: %v", message["message"])
				}
				if errors, ok := message["errors"].([]string); ok {
					if len(errors) == 0 {
						t.Error("Expected validation errors, got empty slice")
					}
				} else {
					t.Error("Expected errors to be a slice of strings")
				}
			} else {
				t.Error("Expected message to be a map with validation details")
			}
		} else {
			t.Error("Expected Echo HTTP error")
		}
	})

	t.Run("Valid struct with no validation tags should pass", func(t *testing.T) {
		type TestStruct struct {
			Name string
		}

		validStruct := TestStruct{
			Name: "",
		}

		err := validator.Validate(validStruct)
		if err != nil {
			t.Errorf("Expected no error for struct with no validation tags, got: %v", err)
		}
	})
}
