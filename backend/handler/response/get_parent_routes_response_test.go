package response

import (
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewParentRoutesResponse(t *testing.T) {
	// Given
	parentRouteName1, _ := vo.NewParentRouteName("テスト路線1")
	parentRouteName2, _ := vo.NewParentRouteName("テスト路線2")

	parentRoute1 := entity.NewParentRoute("route-001", *parentRouteName1)
	parentRoute2 := entity.NewParentRoute("route-002", *parentRouteName2)

	parentRoutes := []entity.ParentRoute{*parentRoute1, *parentRoute2}

	// When
	response := NewParentRoutesResponse(parentRoutes)

	// Then
	if len(response.ParentRoutes) != 2 {
		t.Errorf("Expected 2 parent routes, got %d", len(response.ParentRoutes))
	}

	// Check first parent route
	if response.ParentRoutes[0].ParentRouteId != "route-001" {
		t.Errorf("Expected first ParentRouteId to be 'route-001', got '%s'", response.ParentRoutes[0].ParentRouteId)
	}
	if response.ParentRoutes[0].ParentRouteName != "テスト路線1" {
		t.Errorf("Expected first ParentRouteName to be 'テスト路線1', got '%s'", response.ParentRoutes[0].ParentRouteName)
	}

	// Check second parent route
	if response.ParentRoutes[1].ParentRouteId != "route-002" {
		t.Errorf("Expected second ParentRouteId to be 'route-002', got '%s'", response.ParentRoutes[1].ParentRouteId)
	}
	if response.ParentRoutes[1].ParentRouteName != "テスト路線2" {
		t.Errorf("Expected second ParentRouteName to be 'テスト路線2', got '%s'", response.ParentRoutes[1].ParentRouteName)
	}
}

func TestNewParentRoutesResponseEmpty(t *testing.T) {
	// Given
	parentRoutes := []entity.ParentRoute{}

	// When
	response := NewParentRoutesResponse(parentRoutes)

	// Then
	if len(response.ParentRoutes) != 0 {
		t.Errorf("Expected 0 parent routes, got %d", len(response.ParentRoutes))
	}

	if response.ParentRoutes == nil {
		t.Error("Expected ParentRoutes to be an empty slice, got nil")
	}
}

func TestNewParentRoutesResponseSingleRoute(t *testing.T) {
	// Given
	parentRouteName, _ := vo.NewParentRouteName("単一テスト路線")
	parentRoute := entity.NewParentRoute("single-route", *parentRouteName)
	parentRoutes := []entity.ParentRoute{*parentRoute}

	// When
	response := NewParentRoutesResponse(parentRoutes)

	// Then
	if len(response.ParentRoutes) != 1 {
		t.Errorf("Expected 1 parent route, got %d", len(response.ParentRoutes))
	}

	if response.ParentRoutes[0].ParentRouteId != "single-route" {
		t.Errorf("Expected ParentRouteId to be 'single-route', got '%s'", response.ParentRoutes[0].ParentRouteId)
	}
	if response.ParentRoutes[0].ParentRouteName != "単一テスト路線" {
		t.Errorf("Expected ParentRouteName to be '単一テスト路線', got '%s'", response.ParentRoutes[0].ParentRouteName)
	}
}

func TestNewParentRoutesResponseOrder(t *testing.T) {
	// Given - 順序を確認するため、複数の路線を追加
	parentRouteName1, _ := vo.NewParentRouteName("路線A")
	parentRouteName2, _ := vo.NewParentRouteName("路線B")
	parentRouteName3, _ := vo.NewParentRouteName("路線C")

	parentRoute1 := entity.NewParentRoute("route-a", *parentRouteName1)
	parentRoute2 := entity.NewParentRoute("route-b", *parentRouteName2)
	parentRoute3 := entity.NewParentRoute("route-c", *parentRouteName3)

	parentRoutes := []entity.ParentRoute{*parentRoute1, *parentRoute2, *parentRoute3}

	// When
	response := NewParentRoutesResponse(parentRoutes)

	// Then
	if len(response.ParentRoutes) != 3 {
		t.Errorf("Expected 3 parent routes, got %d", len(response.ParentRoutes))
	}

	// 順序が保持されていることを確認
	expectedRoutes := []struct {
		id   string
		name string
	}{
		{"route-a", "路線A"},
		{"route-b", "路線B"},
		{"route-c", "路線C"},
	}

	for i, expected := range expectedRoutes {
		if response.ParentRoutes[i].ParentRouteId != expected.id {
			t.Errorf("Expected ParentRouteId at index %d to be '%s', got '%s'", i, expected.id, response.ParentRoutes[i].ParentRouteId)
		}
		if response.ParentRoutes[i].ParentRouteName != expected.name {
			t.Errorf("Expected ParentRouteName at index %d to be '%s', got '%s'", i, expected.name, response.ParentRoutes[i].ParentRouteName)
		}
	}
}
