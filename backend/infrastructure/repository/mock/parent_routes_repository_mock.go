package mock

import (
	"context"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

// MockParentRoutesRepository is a mock implementation of ParentRoutesRepository
type MockParentRoutesRepository struct {
	// テスト結果を設定するフィールド
	GetAllResult  []entity.ParentRoute
	GetAllError   error
	GetByIdResult entity.ParentRoute
	GetByIdError  error

	// テスト検証用フィールド
	GetAllCallCount    int
	GetByIdCallCount   int
	LastGetAllContext  context.Context
	LastGetByIdContext context.Context
	LastGetByIdParam   string
}

func (m *MockParentRoutesRepository) GetAll(ctx context.Context) ([]entity.ParentRoute, error) {
	m.GetAllCallCount++
	m.LastGetAllContext = ctx
	return m.GetAllResult, m.GetAllError
}

func (m *MockParentRoutesRepository) GetByParentRouteId(ctx context.Context, parentRouteId string) (entity.ParentRoute, error) {
	m.GetByIdCallCount++
	m.LastGetByIdContext = ctx
	m.LastGetByIdParam = parentRouteId
	return m.GetByIdResult, m.GetByIdError
}

// Reset は各テストの前にモックの状態をリセットするためのヘルパーメソッド
func (m *MockParentRoutesRepository) Reset() {
	m.GetAllResult = nil
	m.GetAllError = nil
	m.GetByIdResult = entity.ParentRoute{}
	m.GetByIdError = nil
	m.GetAllCallCount = 0
	m.GetByIdCallCount = 0
	m.LastGetAllContext = nil
	m.LastGetByIdContext = nil
	m.LastGetByIdParam = ""
}
