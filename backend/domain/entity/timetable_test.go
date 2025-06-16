package entity

import (
	"reflect"
	"testing"
	"time"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewTimetable(t *testing.T) {
	// テスト用の親路線名を作成
	parentRouteNameValue := "東京～横浜"
	parentRouteName, err := vo.NewParentRouteName(parentRouteNameValue)
	if err != nil {
		t.Fatalf("親路線名の作成に失敗しました: %v", err)
	}

	// テスト用のバス停名を作成
	busStopNameValue := "東京駅前"
	busStopName, err := vo.NewBusStopName(busStopNameValue)
	if err != nil {
		t.Fatalf("バス停名の作成に失敗しました: %v", err)
	}

	// テスト用のIDを設定
	parentRouteId := "PR001"
	busStopId := "BS001"

	// テスト用の時刻表エントリーを作成
	entries := createTestTimetableEntries(t)

	// Timetableエンティティを作成
	timetable := NewTimetable(
		parentRouteId,
		*parentRouteName,
		busStopId,
		*busStopName,
		entries,
	)

	// 期待される値と一致するか確認
	if timetable.ParentRouteId != parentRouteId {
		t.Errorf("ParentRouteId = %v, 期待値 %v", timetable.ParentRouteId, parentRouteId)
	}

	if timetable.ParentRouteName.Value() != parentRouteNameValue {
		t.Errorf("ParentRouteName.Value() = %v, 期待値 %v", timetable.ParentRouteName.Value(), parentRouteNameValue)
	}

	if timetable.BusStopId != busStopId {
		t.Errorf("BusStopId = %v, 期待値 %v", timetable.BusStopId, busStopId)
	}

	if timetable.BusStopName.Value() != busStopNameValue {
		t.Errorf("BusStopName.Value() = %v, 期待値 %v", timetable.BusStopName.Value(), busStopNameValue)
	}

	if len(timetable.TimetableEntries) != len(entries) {
		t.Errorf("TimetableEntries length = %v, 期待値 %v", len(timetable.TimetableEntries), len(entries))
	}
}

func TestNewTimetable_WithEmptyEntries(t *testing.T) {
	// テスト用の親路線名を作成
	parentRouteName, _ := vo.NewParentRouteName("テスト路線")
	
	// テスト用のバス停名を作成
	busStopName, _ := vo.NewBusStopName("テストバス停")

	// 空のエントリーリストでテスト
	var entries []TimetableEntry

	// Timetableエンティティを作成
	timetable := NewTimetable("PR002", *parentRouteName, "BS002", *busStopName, entries)

	// エントリーリストが空であることを確認
	if len(timetable.TimetableEntries) != 0 {
		t.Errorf("TimetableEntries length = %v, 期待値 0", len(timetable.TimetableEntries))
	}
}

func TestTimetable_FieldAccess(t *testing.T) {
	// テスト用の親路線名を作成
	parentRouteName, _ := vo.NewParentRouteName("新宿～横浜")
	
	// テスト用のバス停名を作成
	busStopName, _ := vo.NewBusStopName("新宿駅南口")

	// テスト用の時刻表エントリーを作成
	entries := createTestTimetableEntries(t)

	// テスト用のIDを設定
	parentRouteId := "PR003"
	busStopId := "BS003"

	// Timetableエンティティを作成
	timetable := NewTimetable(
		parentRouteId,
		*parentRouteName,
		busStopId,
		*busStopName,
		entries,
	)

	// フィールドへの直接アクセスで値を確認
	if timetable.ParentRouteId != parentRouteId {
		t.Errorf("ParentRouteId = %v, 期待値 %v", timetable.ParentRouteId, parentRouteId)
	}

	if timetable.ParentRouteName.Value() != parentRouteName.Value() {
		t.Errorf("ParentRouteName.Value() = %v, 期待値 %v", timetable.ParentRouteName.Value(), parentRouteName.Value())
	}

	if timetable.BusStopId != busStopId {
		t.Errorf("BusStopId = %v, 期待値 %v", timetable.BusStopId, busStopId)
	}

	if timetable.BusStopName.Value() != busStopName.Value() {
		t.Errorf("BusStopName.Value() = %v, 期待値 %v", timetable.BusStopName.Value(), busStopName.Value())
	}

	// エントリーの検証
	if !reflect.DeepEqual(timetable.TimetableEntries, entries) {
		t.Errorf("TimetableEntries が期待値と一致しません")
	}
}

// テスト用の時刻表エントリー作成ヘルパー関数
func createTestTimetableEntries(t *testing.T) []TimetableEntry {
	// 出発時間の作成
	dt1, err := vo.NewDepartureTime("08:00")
	if err != nil {
		t.Fatalf("出発時間の作成に失敗しました: %v", err)
	}
	dt2, err := vo.NewDepartureTime("09:30")
	if err != nil {
		t.Fatalf("出発時間の作成に失敗しました: %v", err)
	}
	
	// 行先名の作成
	dest1, err := vo.NewDestinationName("横浜駅")
	if err != nil {
		t.Fatalf("行先名の作成に失敗しました: %v", err)
	}
	dest2, err := vo.NewDestinationName("川崎駅")
	if err != nil {
		t.Fatalf("行先名の作成に失敗しました: %v", err)
	}
	
	// 運行日の作成
	mon := vo.NewOperationDay(time.Monday)
	tue := vo.NewOperationDay(time.Tuesday)
	wed := vo.NewOperationDay(time.Wednesday)
	fri := vo.NewOperationDay(time.Friday)
	
	opDays1 := map[vo.OperationDay]struct{}{
		*mon: {},
		*wed: {},
		*fri: {},
	}
	
	opDays2 := map[vo.OperationDay]struct{}{
		*tue: {},
		*fri: {},
	}
	
	// エントリーの作成
	entry1 := NewTimetableEntry(*dt1, opDays1, *dest1)
	entry2 := NewTimetableEntry(*dt2, opDays2, *dest2)
	
	return []TimetableEntry{*entry1, *entry2}
}
