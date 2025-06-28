package response

import (
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewBusStopsResponse(t *testing.T) {
	// Given
	parentRouteName, _ := vo.NewParentRouteName("テスト路線")
	parentRoute := entity.NewParentRoute("route-001", *parentRouteName)

	busStopName1, _ := vo.NewBusStopName("テストバス停1")
	busStopName2, _ := vo.NewBusStopName("テストバス停2")

	busStop1 := entity.NewBusStop("stop-001", *busStopName1)
	busStop2 := entity.NewBusStop("stop-002", *busStopName2)

	busStops := []entity.BusStop{*busStop1, *busStop2}

	// When
	response := NewBusStopsResponse(busStops, *parentRoute)

	// Then
	if response.ParentRouteId != "route-001" {
		t.Errorf("Expected ParentRouteId to be 'route-001', got '%s'", response.ParentRouteId)
	}

	if response.ParentRouteName != "テスト路線" {
		t.Errorf("Expected ParentRouteName to be 'テスト路線', got '%s'", response.ParentRouteName)
	}

	if len(response.BusStops) != 2 {
		t.Errorf("Expected 2 bus stops, got %d", len(response.BusStops))
	}

	// Check first bus stop
	if response.BusStops[0].BusStopId != "stop-001" {
		t.Errorf("Expected first BusStopId to be 'stop-001', got '%s'", response.BusStops[0].BusStopId)
	}
	if response.BusStops[0].BusStopName != "テストバス停1" {
		t.Errorf("Expected first BusStopName to be 'テストバス停1', got '%s'", response.BusStops[0].BusStopName)
	}

	// Check second bus stop
	if response.BusStops[1].BusStopId != "stop-002" {
		t.Errorf("Expected second BusStopId to be 'stop-002', got '%s'", response.BusStops[1].BusStopId)
	}
	if response.BusStops[1].BusStopName != "テストバス停2" {
		t.Errorf("Expected second BusStopName to be 'テストバス停2', got '%s'", response.BusStops[1].BusStopName)
	}
}

func TestNewBusStopsResponseEmptyBusStops(t *testing.T) {
	// Given
	parentRouteName, _ := vo.NewParentRouteName("テスト路線")
	parentRoute := entity.NewParentRoute("route-001", *parentRouteName)
	busStops := []entity.BusStop{}

	// When
	response := NewBusStopsResponse(busStops, *parentRoute)

	// Then
	if response.ParentRouteId != "route-001" {
		t.Errorf("Expected ParentRouteId to be 'route-001', got '%s'", response.ParentRouteId)
	}

	if response.ParentRouteName != "テスト路線" {
		t.Errorf("Expected ParentRouteName to be 'テスト路線', got '%s'", response.ParentRouteName)
	}

	if len(response.BusStops) != 0 {
		t.Errorf("Expected 0 bus stops, got %d", len(response.BusStops))
	}

	if response.BusStops == nil {
		t.Error("Expected BusStops to be an empty slice, got nil")
	}
}

func TestNewBusStopsResponseSingleBusStop(t *testing.T) {
	// Given
	parentRouteName, _ := vo.NewParentRouteName("単一路線")
	parentRoute := entity.NewParentRoute("single-route", *parentRouteName)

	busStopName, _ := vo.NewBusStopName("単一バス停")
	busStop := entity.NewBusStop("single-stop", *busStopName)
	busStops := []entity.BusStop{*busStop}

	// When
	response := NewBusStopsResponse(busStops, *parentRoute)

	// Then
	if response.ParentRouteId != "single-route" {
		t.Errorf("Expected ParentRouteId to be 'single-route', got '%s'", response.ParentRouteId)
	}

	if response.ParentRouteName != "単一路線" {
		t.Errorf("Expected ParentRouteName to be '単一路線', got '%s'", response.ParentRouteName)
	}

	if len(response.BusStops) != 1 {
		t.Errorf("Expected 1 bus stop, got %d", len(response.BusStops))
	}

	if response.BusStops[0].BusStopId != "single-stop" {
		t.Errorf("Expected BusStopId to be 'single-stop', got '%s'", response.BusStops[0].BusStopId)
	}
	if response.BusStops[0].BusStopName != "単一バス停" {
		t.Errorf("Expected BusStopName to be '単一バス停', got '%s'", response.BusStops[0].BusStopName)
	}
}

func TestNewBusStopsResponseOrder(t *testing.T) {
	// Given - 順序を確認するため、複数のバス停を追加
	parentRouteName, _ := vo.NewParentRouteName("順序テスト路線")
	parentRoute := entity.NewParentRoute("order-route", *parentRouteName)

	busStopName1, _ := vo.NewBusStopName("バス停A")
	busStopName2, _ := vo.NewBusStopName("バス停B")
	busStopName3, _ := vo.NewBusStopName("バス停C")

	busStop1 := entity.NewBusStop("stop-a", *busStopName1)
	busStop2 := entity.NewBusStop("stop-b", *busStopName2)
	busStop3 := entity.NewBusStop("stop-c", *busStopName3)

	busStops := []entity.BusStop{*busStop1, *busStop2, *busStop3}

	// When
	response := NewBusStopsResponse(busStops, *parentRoute)

	// Then
	if len(response.BusStops) != 3 {
		t.Errorf("Expected 3 bus stops, got %d", len(response.BusStops))
	}

	// 順序が保持されていることを確認
	expectedStops := []struct {
		id   string
		name string
	}{
		{"stop-a", "バス停A"},
		{"stop-b", "バス停B"},
		{"stop-c", "バス停C"},
	}

	for i, expected := range expectedStops {
		if response.BusStops[i].BusStopId != expected.id {
			t.Errorf("Expected BusStopId at index %d to be '%s', got '%s'", i, expected.id, response.BusStops[i].BusStopId)
		}
		if response.BusStops[i].BusStopName != expected.name {
			t.Errorf("Expected BusStopName at index %d to be '%s', got '%s'", i, expected.name, response.BusStops[i].BusStopName)
		}
	}
}

func TestNewBusStopsResponseParentRouteInfo(t *testing.T) {
	// Given - 親路線の情報が正しく設定されることを確認
	parentRouteName, _ := vo.NewParentRouteName("親路線情報テスト")
	parentRoute := entity.NewParentRoute("parent-info-test", *parentRouteName)

	busStopName, _ := vo.NewBusStopName("テストバス停")
	busStop := entity.NewBusStop("test-stop", *busStopName)
	busStops := []entity.BusStop{*busStop}

	// When
	response := NewBusStopsResponse(busStops, *parentRoute)

	// Then - 親路線情報が正確に反映されていることを確認
	if response.ParentRouteId != parentRoute.ParentRouteId {
		t.Errorf("Expected ParentRouteId to match parent route ID '%s', got '%s'", parentRoute.ParentRouteId, response.ParentRouteId)
	}

	if response.ParentRouteName != parentRoute.ParentRouteName.Value() {
		t.Errorf("Expected ParentRouteName to match parent route name '%s', got '%s'", parentRoute.ParentRouteName.Value(), response.ParentRouteName)
	}
}
