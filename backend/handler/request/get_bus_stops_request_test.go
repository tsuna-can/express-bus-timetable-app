package request

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// TestValidator implements echo.Validator interface for testing
type TestValidator struct {
	validator *validator.Validate
}

func (tv *TestValidator) Validate(i interface{}) error {
	return tv.validator.Struct(i)
}

func TestBusStopsRequest(t *testing.T) {
	e := echo.New()
	e.Validator = &TestValidator{validator: validator.New()}

	t.Run("Valid request should pass", func(t *testing.T) {
		q := make(url.Values)
		q.Set("parent-route-id", "test-route-id")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		c := e.NewContext(req, httptest.NewRecorder())

		busStopsReq, err := NewBusStopsRequest(c)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if busStopsReq.ParentRouteId != "test-route-id" {
			t.Errorf("Expected ParentRouteId to be 'test-route-id', got: %s", busStopsReq.ParentRouteId)
		}
	})

	t.Run("Missing required field should fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, httptest.NewRecorder())

		_, err := NewBusStopsRequest(c)
		if err == nil {
			t.Error("Expected validation error, got nil")
		}
	})
}
