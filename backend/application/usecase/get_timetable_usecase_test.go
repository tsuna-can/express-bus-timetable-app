package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
	"github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure/repository/mock"
)

func TestNewGetTimetableUsecase(t *testing.T) {
	timetableRepo := &mock.MockTimetableRepository{}
	usecase := NewGetTimetableUsecase(timetableRepo)

	if usecase == nil {
		t.Error("usecase should not be nil")
	}

	// 型チェック
	if _, ok := usecase.(*getTimetableUsecase); !ok {
		t.Error("usecase should be of type *getTimetableUsecase")
	}
}

func TestGetTimetableUsecase_GetByParentRouteIdAndBusStopId(t *testing.T) {
	tests := []struct {
		name              string
		parentRouteId     string
		busStopId         string
		setupMock         func(*mock.MockTimetableRepository)
		expectedTimetable entity.Timetable
		expectedError     string
	}{
		{
			name:          "正常系：時刻表データが存在する場合",
			parentRouteId: "route1",
			busStopId:     "stop1",
			setupMock: func(mockRepo *mock.MockTimetableRepository) {
				// テストデータを作成
				parentRouteName, _ := vo.NewParentRouteName("山手線")
				busStopName, _ := vo.NewBusStopName("東京駅")
				departureTime1, _ := vo.NewDepartureTime("08:00")
				departureTime2, _ := vo.NewDepartureTime("08:30")
				destinationName, _ := vo.NewDestinationName("新宿方面")

				// 運行日の作成（平日：月〜金）
				operationDays1 := map[vo.OperationDay]struct{}{
					*vo.NewOperationDay(time.Monday):    {},
					*vo.NewOperationDay(time.Tuesday):   {},
					*vo.NewOperationDay(time.Wednesday): {},
					*vo.NewOperationDay(time.Thursday):  {},
					*vo.NewOperationDay(time.Friday):    {},
				}

				operationDays2 := map[vo.OperationDay]struct{}{
					*vo.NewOperationDay(time.Saturday): {},
					*vo.NewOperationDay(time.Sunday):   {},
				}

				timetableEntries := []entity.TimetableEntry{
					{
						DepartureTime:   *departureTime1,
						OperationDays:   operationDays1,
						DestinationName: *destinationName,
					},
					{
						DepartureTime:   *departureTime2,
						OperationDays:   operationDays2,
						DestinationName: *destinationName,
					},
				}

				expectedTimetable := entity.Timetable{
					ParentRouteId:    "route1",
					ParentRouteName:  *parentRouteName,
					BusStopId:        "stop1",
					BusStopName:      *busStopName,
					TimetableEntries: timetableEntries,
				}

				mockRepo.GetByParentRouteIdAndBusStopIdResult = expectedTimetable
				mockRepo.GetByParentRouteIdAndBusStopIdError = nil
			},
			expectedTimetable: func() entity.Timetable {
				parentRouteName, _ := vo.NewParentRouteName("山手線")
				busStopName, _ := vo.NewBusStopName("東京駅")
				departureTime1, _ := vo.NewDepartureTime("08:00")
				departureTime2, _ := vo.NewDepartureTime("08:30")
				destinationName, _ := vo.NewDestinationName("新宿方面")

				operationDays1 := map[vo.OperationDay]struct{}{
					*vo.NewOperationDay(time.Monday):    {},
					*vo.NewOperationDay(time.Tuesday):   {},
					*vo.NewOperationDay(time.Wednesday): {},
					*vo.NewOperationDay(time.Thursday):  {},
					*vo.NewOperationDay(time.Friday):    {},
				}

				operationDays2 := map[vo.OperationDay]struct{}{
					*vo.NewOperationDay(time.Saturday): {},
					*vo.NewOperationDay(time.Sunday):   {},
				}

				timetableEntries := []entity.TimetableEntry{
					{
						DepartureTime:   *departureTime1,
						OperationDays:   operationDays1,
						DestinationName: *destinationName,
					},
					{
						DepartureTime:   *departureTime2,
						OperationDays:   operationDays2,
						DestinationName: *destinationName,
					},
				}

				return entity.Timetable{
					ParentRouteId:    "route1",
					ParentRouteName:  *parentRouteName,
					BusStopId:        "stop1",
					BusStopName:      *busStopName,
					TimetableEntries: timetableEntries,
				}
			}(),
			expectedError: "",
		},
		{
			name:          "正常系：時刻表エントリが空の場合",
			parentRouteId: "route2",
			busStopId:     "stop2",
			setupMock: func(mockRepo *mock.MockTimetableRepository) {
				parentRouteName, _ := vo.NewParentRouteName("中央線")
				busStopName, _ := vo.NewBusStopName("新宿駅")

				expectedTimetable := entity.Timetable{
					ParentRouteId:    "route2",
					ParentRouteName:  *parentRouteName,
					BusStopId:        "stop2",
					BusStopName:      *busStopName,
					TimetableEntries: []entity.TimetableEntry{},
				}

				mockRepo.GetByParentRouteIdAndBusStopIdResult = expectedTimetable
				mockRepo.GetByParentRouteIdAndBusStopIdError = nil
			},
			expectedTimetable: func() entity.Timetable {
				parentRouteName, _ := vo.NewParentRouteName("中央線")
				busStopName, _ := vo.NewBusStopName("新宿駅")
				return entity.Timetable{
					ParentRouteId:    "route2",
					ParentRouteName:  *parentRouteName,
					BusStopId:        "stop2",
					BusStopName:      *busStopName,
					TimetableEntries: []entity.TimetableEntry{},
				}
			}(),
			expectedError: "",
		},
		{
			name:          "異常系：repositoryでエラーが発生する場合",
			parentRouteId: "route3",
			busStopId:     "stop3",
			setupMock: func(mockRepo *mock.MockTimetableRepository) {
				mockRepo.GetByParentRouteIdAndBusStopIdResult = entity.Timetable{}
				mockRepo.GetByParentRouteIdAndBusStopIdError = errors.New("database connection error")
			},
			expectedTimetable: entity.Timetable{},
			expectedError:     "failed to get timetable for parent route route3 and bus stop stop3: database connection error",
		},
		{
			name:          "異常系：時刻表が見つからない場合",
			parentRouteId: "nonexistent_route",
			busStopId:     "nonexistent_stop",
			setupMock: func(mockRepo *mock.MockTimetableRepository) {
				mockRepo.GetByParentRouteIdAndBusStopIdResult = entity.Timetable{}
				mockRepo.GetByParentRouteIdAndBusStopIdError = errors.New("timetable not found")
			},
			expectedTimetable: entity.Timetable{},
			expectedError:     "failed to get timetable for parent route nonexistent_route and bus stop nonexistent_stop: timetable not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			timetableRepo := &mock.MockTimetableRepository{}
			tt.setupMock(timetableRepo)
			usecase := NewGetTimetableUsecase(timetableRepo)
			ctx := context.Background()

			// Act
			result, err := usecase.GetByParentRouteIdAndBusStopId(ctx, tt.parentRouteId, tt.busStopId)

			// Assert
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if !contains(err.Error(), tt.expectedError) {
					t.Errorf("expected error to contain '%s', but got '%s'", tt.expectedError, err.Error())
				}
				if !isEmptyTimetable(result) {
					t.Errorf("expected timetable to be empty when error occurs, but got %v", result)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				if !equalTimetables(tt.expectedTimetable, result) {
					t.Errorf("expected timetable %v, but got %v", tt.expectedTimetable, result)
				}
			}

			// モックが正しく呼ばれたことを確認
			if timetableRepo.GetByParentRouteIdAndBusStopIdCallCount != 1 {
				t.Errorf("expected TimetableRepository.GetByParentRouteIdAndBusStopId to be called once, but was called %d times", timetableRepo.GetByParentRouteIdAndBusStopIdCallCount)
			}
			if timetableRepo.LastGetByParentRouteIdAndBusStopIdParentRouteParam != tt.parentRouteId {
				t.Errorf("expected parentRouteId to be '%s', but got '%s'", tt.parentRouteId, timetableRepo.LastGetByParentRouteIdAndBusStopIdParentRouteParam)
			}
			if timetableRepo.LastGetByParentRouteIdAndBusStopIdBusStopParam != tt.busStopId {
				t.Errorf("expected busStopId to be '%s', but got '%s'", tt.busStopId, timetableRepo.LastGetByParentRouteIdAndBusStopIdBusStopParam)
			}
		})
	}
}

func TestGetTimetableUsecase_GetByParentRouteIdAndBusStopId_WithContext(t *testing.T) {
	t.Run("contextが正しく渡されることを確認", func(t *testing.T) {
		// Arrange
		timetableRepo := &mock.MockTimetableRepository{}

		// カスタム型でcontextキーを定義
		type contextKey string
		const testKey contextKey = "test_key"
		ctx := context.WithValue(context.Background(), testKey, "test_value")

		// 正常系のレスポンスを設定
		parentRouteName, _ := vo.NewParentRouteName("テスト線")
		busStopName, _ := vo.NewBusStopName("テスト駅")
		timetable := entity.Timetable{
			ParentRouteId:    "test_route",
			ParentRouteName:  *parentRouteName,
			BusStopId:        "test_stop",
			BusStopName:      *busStopName,
			TimetableEntries: []entity.TimetableEntry{},
		}

		timetableRepo.GetByParentRouteIdAndBusStopIdResult = timetable
		timetableRepo.GetByParentRouteIdAndBusStopIdError = nil

		usecase := NewGetTimetableUsecase(timetableRepo)

		// Act
		_, err := usecase.GetByParentRouteIdAndBusStopId(ctx, "test_route", "test_stop")

		// Assert
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		// contextが正しく渡されたことを確認
		if timetableRepo.LastGetByParentRouteIdAndBusStopIdContext != ctx {
			t.Error("context was not passed correctly to TimetableRepository")
		}
	})
}

// ヘルパー関数
func equalTimetables(a, b entity.Timetable) bool {
	if a.ParentRouteId != b.ParentRouteId ||
		a.ParentRouteName.Value() != b.ParentRouteName.Value() ||
		a.BusStopId != b.BusStopId ||
		a.BusStopName.Value() != b.BusStopName.Value() {
		return false
	}

	if len(a.TimetableEntries) != len(b.TimetableEntries) {
		return false
	}

	for i := range a.TimetableEntries {
		if !equalTimetableEntries(a.TimetableEntries[i], b.TimetableEntries[i]) {
			return false
		}
	}

	return true
}

func equalTimetableEntries(a, b entity.TimetableEntry) bool {
	if a.DepartureTime.Value() != b.DepartureTime.Value() ||
		a.DestinationName.Value() != b.DestinationName.Value() {
		return false
	}

	if len(a.OperationDays) != len(b.OperationDays) {
		return false
	}

	for opDay := range a.OperationDays {
		if _, exists := b.OperationDays[opDay]; !exists {
			return false
		}
	}

	return true
}

func isEmptyTimetable(t entity.Timetable) bool {
	return t.ParentRouteId == "" &&
		t.ParentRouteName.Value() == "" &&
		t.BusStopId == "" &&
		t.BusStopName.Value() == "" &&
		len(t.TimetableEntries) == 0
}
