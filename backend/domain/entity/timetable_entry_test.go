package entity

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewTimetableEntry(t *testing.T) {
	// 出発時間の作成
	departureTimeValue := "08:30"
	departureTime, err := vo.NewDepartureTime(departureTimeValue)
	if err != nil {
		t.Fatalf("出発時間の作成に失敗しました: %v", err)
	}

	// 行先名の作成
	destNameValue := "渋谷駅"
	destName, err := vo.NewDestinationName(destNameValue)
	if err != nil {
		t.Fatalf("行先名の作成に失敗しました: %v", err)
	}

	// 運行日の作成
	monday := vo.NewOperationDay(time.Monday)
	wednesday := vo.NewOperationDay(time.Wednesday)
	friday := vo.NewOperationDay(time.Friday)
	
	operationDays := make(map[vo.OperationDay]struct{})
	operationDays[*monday] = struct{}{}
	operationDays[*wednesday] = struct{}{}
	operationDays[*friday] = struct{}{}

	// TimetableEntryエンティティ作成
	entry := NewTimetableEntry(*departureTime, operationDays, *destName)

	// DepartureTimeが正しく設定されているか確認
	if entry.DepartureTime.Value() != departureTimeValue {
		t.Errorf("DepartureTime.Value() = %v, 期待値 %v", entry.DepartureTime.Value(), departureTimeValue)
	}

	// DestinationNameが正しく設定されているか確認
	if entry.DestinationName.Value() != destNameValue {
		t.Errorf("DestinationName.Value() = %v, 期待値 %v", entry.DestinationName.Value(), destNameValue)
	}

	// OperationDaysのサイズ確認
	if len(entry.OperationDays) != 3 {
		t.Errorf("OperationDays length = %v, 期待値 3", len(entry.OperationDays))
	}

	// OperationDaysの内容確認
	_, hasMonday := entry.OperationDays[*monday]
	_, hasWednesday := entry.OperationDays[*wednesday]
	_, hasFriday := entry.OperationDays[*friday]
	
	if !hasMonday || !hasWednesday || !hasFriday {
		t.Errorf("OperationDaysに特定の曜日が含まれていません")
	}
}

func TestTimetableEntry_OperationDaysAsIntSlice(t *testing.T) {
	// セットアップ
	departureTime, _ := vo.NewDepartureTime("09:45")
	destName, _ := vo.NewDestinationName("東京駅")
	
	// 曜日の作成（月、水、日曜）
	monday := vo.NewOperationDay(time.Monday)
	wednesday := vo.NewOperationDay(time.Wednesday)
	sunday := vo.NewOperationDay(time.Sunday)
	
	operationDays := make(map[vo.OperationDay]struct{})
	operationDays[*monday] = struct{}{}
	operationDays[*wednesday] = struct{}{}
	operationDays[*sunday] = struct{}{}

	entry := NewTimetableEntry(*departureTime, operationDays, *destName)

	// メソッドのテスト
	result := entry.OperationDaysAsIntSlice()
	
	// 期待結果: [1, 3, 7] (月曜=1, 水曜=3, 日曜=7)
	expected := []int{1, 3, 7}
	
	// スライスは順不同で返される可能性があるためソートして比較
	sort.Ints(result)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("OperationDaysAsIntSlice() = %v, 期待値 %v", result, expected)
	}
}

func TestTimetableEntry_EmptyOperationDays(t *testing.T) {
	// 出発時間と行先名の作成
	departureTime, _ := vo.NewDepartureTime("12:00")
	destName, _ := vo.NewDestinationName("品川駅")
	
	// 空の運行日マップ
	operationDays := make(map[vo.OperationDay]struct{})

	entry := NewTimetableEntry(*departureTime, operationDays, *destName)

	// 空の運行日リストから変換された結果が空のスライスであることを確認
	result := entry.OperationDaysAsIntSlice()
	
	if len(result) != 0 {
		t.Errorf("空の運行日マップから変換された結果が空のスライスではありません。結果: %v", result)
	}
}
