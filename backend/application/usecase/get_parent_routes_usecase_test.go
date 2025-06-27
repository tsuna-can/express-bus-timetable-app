package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/mock"
)

func TestNewGetParentRoutesUsecase(t *testing.T) {
	mockRepo := &mock.MockParentRoutesRepository{}
	usecase := NewGetParentRoutesUsecase(mockRepo)

	if usecase == nil {
		t.Error("usecase should not be nil")
	}

	// 型チェック
	if _, ok := usecase.(*getParentRoutesUsecase); !ok {
		t.Error("usecase should be of type *getParentRoutesUsecase")
	}
}

func TestGetParentRoutesUsecase_GetAll(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*mock.MockParentRoutesRepository)
		expectedRoutes []entity.ParentRoute
		expectedError  string
	}{
		{
			name: "正常系：複数のparent routeが存在する場合",
			setupMock: func(mockRepo *mock.MockParentRoutesRepository) {
				parentRouteName1, _ := vo.NewParentRouteName("東京駅行き")
				parentRouteName2, _ := vo.NewParentRouteName("新宿駅行き")

				mockRepo.GetAllResult = []entity.ParentRoute{
					{ParentRouteId: "route1", ParentRouteName: *parentRouteName1},
					{ParentRouteId: "route2", ParentRouteName: *parentRouteName2},
				}
				mockRepo.GetAllError = nil
			},
			expectedRoutes: func() []entity.ParentRoute {
				parentRouteName1, _ := vo.NewParentRouteName("東京駅行き")
				parentRouteName2, _ := vo.NewParentRouteName("新宿駅行き")
				return []entity.ParentRoute{
					{ParentRouteId: "route1", ParentRouteName: *parentRouteName1},
					{ParentRouteId: "route2", ParentRouteName: *parentRouteName2},
				}
			}(),
			expectedError: "",
		},
		{
			name: "正常系：parent routeが存在しない場合（空のスライス）",
			setupMock: func(mockRepo *mock.MockParentRoutesRepository) {
				mockRepo.GetAllResult = []entity.ParentRoute{}
				mockRepo.GetAllError = nil
			},
			expectedRoutes: []entity.ParentRoute{},
			expectedError:  "",
		},
		{
			name: "異常系：repositoryでエラーが発生する場合",
			setupMock: func(mockRepo *mock.MockParentRoutesRepository) {
				mockRepo.GetAllResult = []entity.ParentRoute{}
				mockRepo.GetAllError = errors.New("database connection error")
			},
			expectedRoutes: nil,
			expectedError:  "failed to get all parent routes: database connection error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := &mock.MockParentRoutesRepository{}
			tt.setupMock(mockRepo)
			usecase := NewGetParentRoutesUsecase(mockRepo)
			ctx := context.Background()

			// Act
			result, err := usecase.GetAll(ctx)

			// Assert
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if !contains(err.Error(), tt.expectedError) {
					t.Errorf("expected error to contain '%s', but got '%s'", tt.expectedError, err.Error())
				}
				if result != nil {
					t.Errorf("expected result to be nil when error occurs, but got %v", result)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				if !equalSlices(tt.expectedRoutes, result) {
					t.Errorf("expected %v, but got %v", tt.expectedRoutes, result)
				}
			}

			// モックが呼ばれたことを確認
			if mockRepo.GetAllCallCount != 1 {
				t.Errorf("expected GetAll to be called once, but was called %d times", mockRepo.GetAllCallCount)
			}
		})
	}
}

func TestGetParentRoutesUsecase_GetAll_WithContext(t *testing.T) {
	t.Run("contextが正しく渡されることを確認", func(t *testing.T) {
		// Arrange
		mockRepo := &mock.MockParentRoutesRepository{}

		// カスタム型でcontextキーを定義
		type contextKey string
		const testKey contextKey = "test_key"
		ctx := context.WithValue(context.Background(), testKey, "test_value")

		mockRepo.GetAllResult = []entity.ParentRoute{}
		mockRepo.GetAllError = nil
		usecase := NewGetParentRoutesUsecase(mockRepo)

		// Act
		_, err := usecase.GetAll(ctx)

		// Assert
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		// contextが正しく渡されたことを確認
		if mockRepo.LastGetAllContext != ctx {
			t.Error("context was not passed correctly to repository")
		}

		if mockRepo.GetAllCallCount != 1 {
			t.Errorf("expected GetAll to be called once, but was called %d times", mockRepo.GetAllCallCount)
		}
	})
}

func equalSlices(a, b []entity.ParentRoute) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].ParentRouteId != b[i].ParentRouteId ||
			a[i].ParentRouteName.Value() != b[i].ParentRouteName.Value() {
			return false
		}
	}
	return true
}
