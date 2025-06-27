package mock

import (
	"context"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

// MockBusStopsRepository is a mock implementation of BusStopsRepository
type MockBusStopsRepository struct {
	// テスト結果を設定するフィールド
	GetByParentRouteIdResult []entity.BusStop
	GetByParentRouteIdError  error

	// テスト検証用フィールド
	GetByParentRouteIdCallCount   int
	LastGetByParentRouteIdContext context.Context
	LastGetByParentRouteIdParam   string
}

func (m *MockBusStopsRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) ([]entity.BusStop, error) {
	m.GetByParentRouteIdCallCount++
	m.LastGetByParentRouteIdContext = ctx
	m.LastGetByParentRouteIdParam = parentRouteId
	return m.GetByParentRouteIdResult, m.GetByParentRouteIdError
}

// Reset は各テストの前にモックの状態をリセットするためのヘルパーメソッド
func (m *MockBusStopsRepository) Reset() {
	m.GetByParentRouteIdResult = nil
	m.GetByParentRouteIdError = nil
	m.GetByParentRouteIdCallCount = 0
	m.LastGetByParentRouteIdContext = nil
	m.LastGetByParentRouteIdParam = ""
}
