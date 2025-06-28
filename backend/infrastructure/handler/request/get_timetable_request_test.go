package request

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func TestTimetableRequest(t *testing.T) {
	e := echo.New()
	e.Validator = &TestValidator{validator: validator.New()}

	t.Run("Valid request should pass", func(t *testing.T) {
		q := make(url.Values)
		q.Set("parent-route-id", "test-parent-route-id")
		q.Set("bus-stop-id", "test-bus-stop-id")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		c := e.NewContext(req, httptest.NewRecorder())

		timetableReq, err := NewTimetableRequest(c)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if timetableReq.ParentRouteId != "test-parent-route-id" {
			t.Errorf("Expected ParentRouteId to be 'test-parent-route-id', got: %s", timetableReq.ParentRouteId)
		}

		if timetableReq.BusStopId != "test-bus-stop-id" {
			t.Errorf("Expected BusStopId to be 'test-bus-stop-id', got: %s", timetableReq.BusStopId)
		}
	})

	t.Run("Missing parent-route-id should fail", func(t *testing.T) {
		q := make(url.Values)
		q.Set("bus-stop-id", "test-bus-stop-id")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		c := e.NewContext(req, httptest.NewRecorder())

		_, err := NewTimetableRequest(c)
		if err == nil {
			t.Error("Expected validation error, got nil")
		}
	})

	t.Run("Missing bus-stop-id should fail", func(t *testing.T) {
		q := make(url.Values)
		q.Set("parent-route-id", "test-parent-route-id")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		c := e.NewContext(req, httptest.NewRecorder())

		_, err := NewTimetableRequest(c)
		if err == nil {
			t.Error("Expected validation error, got nil")
		}
	})

	t.Run("Missing both parameters should fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, httptest.NewRecorder())

		_, err := NewTimetableRequest(c)
		if err == nil {
			t.Error("Expected validation error, got nil")
		}
	})
}
