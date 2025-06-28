package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/mock"
)

func TestNewGetBusStopsUsecase(t *testing.T) {
	busStopsRepo := &mock.MockBusStopsRepository{}
	parentRoutesRepo := &mock.MockParentRoutesRepository{}
	usecase := NewGetBusStopsUsecase(busStopsRepo, parentRoutesRepo)

	if usecase == nil {
		t.Error("usecase should not be nil")
	}

	// 型チェック
	if _, ok := usecase.(*getBusStopsUsecase); !ok {
		t.Error("usecase should be of type *getBusStopsUsecase")
	}
}

func TestGetBusStopsUsecase_GetByParentRouteId(t *testing.T) {
	tests := []struct {
		name                 string
		parentRouteId        string
		setupBusStopsMock    func(*mock.MockBusStopsRepository)
		setupParentRouteMock func(*mock.MockParentRoutesRepository)
		expectedBusStops     []entity.BusStop
		expectedParentRoute  entity.ParentRoute
		expectedError        string
	}{
		{
			name:          "正常系：バス停とparent routeが存在する場合",
			parentRouteId: "route1",
			setupBusStopsMock: func(mockRepo *mock.MockBusStopsRepository) {
				busStopName1, _ := vo.NewBusStopName("東京駅")
				busStopName2, _ := vo.NewBusStopName("新宿駅")

				mockRepo.GetByParentRouteIdResult = []entity.BusStop{
					{BusStopId: "stop1", BusStopName: *busStopName1},
					{BusStopId: "stop2", BusStopName: *busStopName2},
				}
				mockRepo.GetByParentRouteIdError = nil
			},
			setupParentRouteMock: func(mockRepo *mock.MockParentRoutesRepository) {
				parentRouteName, _ := vo.NewParentRouteName("山手線")
				mockRepo.GetByIdResult = entity.ParentRoute{
					ParentRouteId:   "route1",
					ParentRouteName: *parentRouteName,
				}
				mockRepo.GetByIdError = nil
			},
			expectedBusStops: func() []entity.BusStop {
				busStopName1, _ := vo.NewBusStopName("東京駅")
				busStopName2, _ := vo.NewBusStopName("新宿駅")
				return []entity.BusStop{
					{BusStopId: "stop1", BusStopName: *busStopName1},
					{BusStopId: "stop2", BusStopName: *busStopName2},
				}
			}(),
			expectedParentRoute: func() entity.ParentRoute {
				parentRouteName, _ := vo.NewParentRouteName("山手線")
				return entity.ParentRoute{
					ParentRouteId:   "route1",
					ParentRouteName: *parentRouteName,
				}
			}(),
			expectedError: "",
		},
		{
			name:          "正常系：バス停が存在しない場合（空のスライス）",
			parentRouteId: "route2",
			setupBusStopsMock: func(mockRepo *mock.MockBusStopsRepository) {
				mockRepo.GetByParentRouteIdResult = []entity.BusStop{}
				mockRepo.GetByParentRouteIdError = nil
			},
			setupParentRouteMock: func(mockRepo *mock.MockParentRoutesRepository) {
				parentRouteName, _ := vo.NewParentRouteName("中央線")
				mockRepo.GetByIdResult = entity.ParentRoute{
					ParentRouteId:   "route2",
					ParentRouteName: *parentRouteName,
				}
				mockRepo.GetByIdError = nil
			},
			expectedBusStops: []entity.BusStop{},
			expectedParentRoute: func() entity.ParentRoute {
				parentRouteName, _ := vo.NewParentRouteName("中央線")
				return entity.ParentRoute{
					ParentRouteId:   "route2",
					ParentRouteName: *parentRouteName,
				}
			}(),
			expectedError: "",
		},
		{
			name:          "異常系：バス停取得でエラーが発生する場合",
			parentRouteId: "route3",
			setupBusStopsMock: func(mockRepo *mock.MockBusStopsRepository) {
				mockRepo.GetByParentRouteIdResult = []entity.BusStop{}
				mockRepo.GetByParentRouteIdError = errors.New("database connection error")
			},
			setupParentRouteMock: func(mockRepo *mock.MockParentRoutesRepository) {
				// バス停取得でエラーが発生するので、parent route取得は呼ばれない
			},
			expectedBusStops:    nil,
			expectedParentRoute: entity.ParentRoute{},
			expectedError:       "failed to get bus stops for parent route route3: database connection error",
		},
		{
			name:          "異常系：parent route取得でエラーが発生する場合",
			parentRouteId: "route4",
			setupBusStopsMock: func(mockRepo *mock.MockBusStopsRepository) {
				busStopName, _ := vo.NewBusStopName("渋谷駅")
				mockRepo.GetByParentRouteIdResult = []entity.BusStop{
					{BusStopId: "stop1", BusStopName: *busStopName},
				}
				mockRepo.GetByParentRouteIdError = nil
			},
			setupParentRouteMock: func(mockRepo *mock.MockParentRoutesRepository) {
				mockRepo.GetByIdResult = entity.ParentRoute{}
				mockRepo.GetByIdError = errors.New("parent route not found")
			},
			expectedBusStops:    nil,
			expectedParentRoute: entity.ParentRoute{},
			expectedError:       "failed to get parent route route4: parent route not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			busStopsRepo := &mock.MockBusStopsRepository{}
			parentRoutesRepo := &mock.MockParentRoutesRepository{}

			tt.setupBusStopsMock(busStopsRepo)
			tt.setupParentRouteMock(parentRoutesRepo)

			usecase := NewGetBusStopsUsecase(busStopsRepo, parentRoutesRepo)
			ctx := context.Background()

			// Act
			resultBusStops, resultParentRoute, err := usecase.GetByParentRouteId(ctx, tt.parentRouteId)

			// Assert
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if !contains(err.Error(), tt.expectedError) {
					t.Errorf("expected error to contain '%s', but got '%s'", tt.expectedError, err.Error())
				}
				if resultBusStops != nil {
					t.Errorf("expected bus stops to be nil when error occurs, but got %v", resultBusStops)
				}
				if !isEmptyParentRoute(resultParentRoute) {
					t.Errorf("expected parent route to be empty when error occurs, but got %v", resultParentRoute)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				if !equalBusStopsSlices(tt.expectedBusStops, resultBusStops) {
					t.Errorf("expected bus stops %v, but got %v", tt.expectedBusStops, resultBusStops)
				}
				if !equalParentRoutes(tt.expectedParentRoute, resultParentRoute) {
					t.Errorf("expected parent route %v, but got %v", tt.expectedParentRoute, resultParentRoute)
				}
			}

			// モックが正しく呼ばれたことを確認
			if busStopsRepo.GetByParentRouteIdCallCount != 1 {
				t.Errorf("expected BusStopsRepository.GetByParentRouteId to be called once, but was called %d times", busStopsRepo.GetByParentRouteIdCallCount)
			}
			if busStopsRepo.LastGetByParentRouteIdParam != tt.parentRouteId {
				t.Errorf("expected parentRouteId to be '%s', but got '%s'", tt.parentRouteId, busStopsRepo.LastGetByParentRouteIdParam)
			}

			// バス停取得でエラーが発生しない場合のみ、parent route取得が呼ばれる
			if tt.expectedError == "" || !contains(tt.expectedError, "failed to get bus stops") {
				if parentRoutesRepo.GetByIdCallCount != 1 {
					t.Errorf("expected ParentRoutesRepository.GetByParentRouteId to be called once, but was called %d times", parentRoutesRepo.GetByIdCallCount)
				}
				if parentRoutesRepo.LastGetByIdParam != tt.parentRouteId {
					t.Errorf("expected parentRouteId to be '%s', but got '%s'", tt.parentRouteId, parentRoutesRepo.LastGetByIdParam)
				}
			}
		})
	}
}

func TestGetBusStopsUsecase_GetByParentRouteId_WithContext(t *testing.T) {
	t.Run("contextが正しく渡されることを確認", func(t *testing.T) {
		// Arrange
		busStopsRepo := &mock.MockBusStopsRepository{}
		parentRoutesRepo := &mock.MockParentRoutesRepository{}

		// カスタム型でcontextキーを定義
		type contextKey string
		const testKey contextKey = "test_key"
		ctx := context.WithValue(context.Background(), testKey, "test_value")

		// 正常系のレスポンスを設定
		busStopName, _ := vo.NewBusStopName("テスト駅")
		busStopsRepo.GetByParentRouteIdResult = []entity.BusStop{
			{BusStopId: "stop1", BusStopName: *busStopName},
		}
		busStopsRepo.GetByParentRouteIdError = nil

		parentRouteName, _ := vo.NewParentRouteName("テスト線")
		parentRoutesRepo.GetByIdResult = entity.ParentRoute{
			ParentRouteId:   "route1",
			ParentRouteName: *parentRouteName,
		}
		parentRoutesRepo.GetByIdError = nil

		usecase := NewGetBusStopsUsecase(busStopsRepo, parentRoutesRepo)

		// Act
		_, _, err := usecase.GetByParentRouteId(ctx, "route1")

		// Assert
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		// contextが正しく渡されたことを確認
		if busStopsRepo.LastGetByParentRouteIdContext != ctx {
			t.Error("context was not passed correctly to BusStopsRepository")
		}
		if parentRoutesRepo.LastGetByIdContext != ctx {
			t.Error("context was not passed correctly to ParentRoutesRepository")
		}
	})
}

// ヘルパー関数
func equalBusStopsSlices(a, b []entity.BusStop) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].BusStopId != b[i].BusStopId ||
			a[i].BusStopName.Value() != b[i].BusStopName.Value() {
			return false
		}
	}
	return true
}

func equalParentRoutes(a, b entity.ParentRoute) bool {
	return a.ParentRouteId == b.ParentRouteId &&
		a.ParentRouteName.Value() == b.ParentRouteName.Value()
}

func isEmptyParentRoute(pr entity.ParentRoute) bool {
	return pr.ParentRouteId == "" && pr.ParentRouteName.Value() == ""
}
