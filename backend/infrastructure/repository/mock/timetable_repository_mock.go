package mock

import (
	"context"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

// MockTimetableRepository is a mock implementation of TimetableRepository
type MockTimetableRepository struct {
	// テスト結果を設定するフィールド
	GetByParentRouteIdAndBusStopIdResult entity.Timetable
	GetByParentRouteIdAndBusStopIdError  error

	// テスト検証用フィールド
	GetByParentRouteIdAndBusStopIdCallCount            int
	LastGetByParentRouteIdAndBusStopIdContext          context.Context
	LastGetByParentRouteIdAndBusStopIdParentRouteParam string
	LastGetByParentRouteIdAndBusStopIdBusStopParam     string
}

func (m *MockTimetableRepository) GetByParentRouteIdAndBusStopId(ctx context.Context, parentRouteId string, busStopId string) (entity.Timetable, error) {
	m.GetByParentRouteIdAndBusStopIdCallCount++
	m.LastGetByParentRouteIdAndBusStopIdContext = ctx
	m.LastGetByParentRouteIdAndBusStopIdParentRouteParam = parentRouteId
	m.LastGetByParentRouteIdAndBusStopIdBusStopParam = busStopId
	return m.GetByParentRouteIdAndBusStopIdResult, m.GetByParentRouteIdAndBusStopIdError
}

// Reset は各テストの前にモックの状態をリセットするためのヘルパーメソッド
func (m *MockTimetableRepository) Reset() {
	m.GetByParentRouteIdAndBusStopIdResult = entity.Timetable{}
	m.GetByParentRouteIdAndBusStopIdError = nil
	m.GetByParentRouteIdAndBusStopIdCallCount = 0
	m.LastGetByParentRouteIdAndBusStopIdContext = nil
	m.LastGetByParentRouteIdAndBusStopIdParentRouteParam = ""
	m.LastGetByParentRouteIdAndBusStopIdBusStopParam = ""
}
