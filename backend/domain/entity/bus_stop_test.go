package entity

import (
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewBusStop(t *testing.T) {
	// テスト用のバス停名を作成
	busStopNameValue := "東京駅バスターミナル"
	busStopName, err := vo.NewBusStopName(busStopNameValue)
	if err != nil {
		t.Fatalf("バス停名の作成に失敗しました: %v", err)
	}

	// テスト用のIDを設定
	busStopId := "BS001"

	// BusStopエンティティを作成
	busStop := NewBusStop(busStopId, *busStopName)

	// 期待される値と一致するか確認
	if busStop.BusStopId != busStopId {
		t.Errorf("BusStopId = %v, 期待値 %v", busStop.BusStopId, busStopId)
	}

	if busStop.BusStopName.Value() != busStopNameValue {
		t.Errorf("BusStopName.Value() = %v, 期待値 %v", busStop.BusStopName.Value(), busStopNameValue)
	}
}

func TestNewBusStop_WithEmptyId(t *testing.T) {
	// 空のバス停IDでテスト
	busStopId := ""

	// テスト用のバス停名を作成
	busStopName, err := vo.NewBusStopName("有効なバス停名")
	if err != nil {
		t.Fatalf("バス停名の作成に失敗しました: %v", err)
	}

	// 空のIDでもエンティティは作成される（現在の実装ではIDのバリデーションはない）
	busStop := NewBusStop(busStopId, *busStopName)

	// IDが空文字であることを確認
	if busStop.BusStopId != busStopId {
		t.Errorf("BusStopId = %v, 期待値は空文字", busStop.BusStopId)
	}
}

func TestBusStop_FieldAccess(t *testing.T) {
	// テストデータ準備
	busStopId := "BS002"
	busStopNameValue := "渋谷駅"
	busStopName, err := vo.NewBusStopName(busStopNameValue)
	if err != nil {
		t.Fatalf("バス停名の作成に失敗しました: %v", err)
	}

	// エンティティ作成
	busStop := NewBusStop(busStopId, *busStopName)

	// フィールドへの直接アクセスで値を確認
	if busStop.BusStopId != busStopId {
		t.Errorf("BusStopId = %v, 期待値 %v", busStop.BusStopId, busStopId)
	}

	if busStop.BusStopName.Value() != busStopNameValue {
		t.Errorf("BusStopName.Value() = %v, 期待値 %v", busStop.BusStopName.Value(), busStopNameValue)
	}
}
