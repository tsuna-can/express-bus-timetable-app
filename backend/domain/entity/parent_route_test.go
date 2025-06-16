package entity

import (
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewParentRoute(t *testing.T) {
	// テスト用の親路線名を作成
	routeNameValue := "東京～大阪線"
	parentRouteName, err := vo.NewParentRouteName(routeNameValue)
	if err != nil {
		t.Fatalf("親路線名の作成に失敗しました: %v", err)
	}

	// テスト用のIDを設定
	parentRouteId := "PR001"

	// ParentRouteエンティティを作成
	parentRoute := NewParentRoute(parentRouteId, *parentRouteName)

	// 期待される値と一致するか確認
	if parentRoute.ParentRouteId != parentRouteId {
		t.Errorf("ParentRouteId = %v, 期待値 %v", parentRoute.ParentRouteId, parentRouteId)
	}

	if parentRoute.ParentRouteName.Value() != routeNameValue {
		t.Errorf("ParentRouteName.Value() = %v, 期待値 %v", parentRoute.ParentRouteName.Value(), routeNameValue)
	}
}

func TestNewParentRoute_WithEmptyId(t *testing.T) {
	// 空の親路線IDでテスト
	parentRouteId := ""
	
	// テスト用の親路線名を作成
	parentRouteName, err := vo.NewParentRouteName("有効な親路線名")
	if err != nil {
		t.Fatalf("親路線名の作成に失敗しました: %v", err)
	}

	// 空のIDでもエンティティは作成される（現在の実装ではIDのバリデーションはない）
	parentRoute := NewParentRoute(parentRouteId, *parentRouteName)

	// IDが空文字であることを確認
	if parentRoute.ParentRouteId != parentRouteId {
		t.Errorf("ParentRouteId = %v, 期待値は空文字", parentRoute.ParentRouteId)
	}
}

func TestParentRoute_FieldAccess(t *testing.T) {
	// テストデータ準備
	parentRouteId := "PR002"
	routeNameValue := "東京～名古屋線"
	parentRouteName, err := vo.NewParentRouteName(routeNameValue)
	if err != nil {
		t.Fatalf("親路線名の作成に失敗しました: %v", err)
	}

	// エンティティ作成
	parentRoute := NewParentRoute(parentRouteId, *parentRouteName)

	// フィールドへの直接アクセスで値を確認
	if parentRoute.ParentRouteId != parentRouteId {
		t.Errorf("ParentRouteId = %v, 期待値 %v", parentRoute.ParentRouteId, parentRouteId)
	}

	if parentRoute.ParentRouteName.Value() != routeNameValue {
		t.Errorf("ParentRouteName.Value() = %v, 期待値 %v", parentRoute.ParentRouteName.Value(), routeNameValue)
	}
}
