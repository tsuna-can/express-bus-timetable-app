package factory

import (
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewParentRouteFactory(t *testing.T) {
	factory := NewParentRouteFactory()
	if factory == nil {
		t.Error("NewParentRouteFactory() should not return nil")
	}
}

func TestParentRouteFactory_ReconstructFromRawData(t *testing.T) {
	factory := NewParentRouteFactory()

	tests := []struct {
		name    string
		rawData ParentRouteRawData
		want    *entity.ParentRoute
		wantErr bool
	}{
		{
			name: "正常系：有効なデータ",
			rawData: ParentRouteRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "山手線",
			},
			want: func() *entity.ParentRoute {
				name, _ := vo.NewParentRouteName("山手線")
				return entity.NewParentRoute("route001", *name)
			}(),
			wantErr: false,
		},
		{
			name: "異常系：空の路線名",
			rawData: ParentRouteRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系：空白のみの路線名",
			rawData: ParentRouteRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "   ",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := factory.ReconstructFromRawData(tt.rawData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReconstructFromRawData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Error("ReconstructFromRawData() should not return nil when no error expected")
					return
				}
				if got.ParentRouteId != tt.want.ParentRouteId {
					t.Errorf("ReconstructFromRawData() ParentRouteId = %v, want %v", got.ParentRouteId, tt.want.ParentRouteId)
				}
				if got.ParentRouteName.Value() != tt.want.ParentRouteName.Value() {
					t.Errorf("ReconstructFromRawData() ParentRouteName = %v, want %v", got.ParentRouteName.Value(), tt.want.ParentRouteName.Value())
				}
			}
		})
	}
}

func TestParentRouteFactory_ReconstructManyFromRawData(t *testing.T) {
	factory := NewParentRouteFactory()

	tests := []struct {
		name     string
		rawData  []ParentRouteRawData
		wantLen  int
		wantErr  bool
		errIndex int // エラーが発生する予想インデックス
	}{
		{
			name: "正常系：複数の有効なデータ",
			rawData: []ParentRouteRawData{
				{ParentRouteId: "route001", ParentRouteName: "山手線"},
				{ParentRouteId: "route002", ParentRouteName: "中央線"},
				{ParentRouteId: "route003", ParentRouteName: "京浜東北線"},
			},
			wantLen: 3,
			wantErr: false,
		},
		{
			name:    "正常系：空のスライス",
			rawData: []ParentRouteRawData{},
			wantLen: 0,
			wantErr: false,
		},
		{
			name: "異常系：2番目のデータが無効",
			rawData: []ParentRouteRawData{
				{ParentRouteId: "route001", ParentRouteName: "山手線"},
				{ParentRouteId: "route002", ParentRouteName: ""}, // 無効
				{ParentRouteId: "route003", ParentRouteName: "京浜東北線"},
			},
			wantLen:  0,
			wantErr:  true,
			errIndex: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := factory.ReconstructManyFromRawData(tt.rawData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReconstructManyFromRawData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != tt.wantLen {
					t.Errorf("ReconstructManyFromRawData() length = %v, want %v", len(got), tt.wantLen)
				}
				// 各要素の内容も確認
				for i, parentRoute := range got {
					expectedId := tt.rawData[i].ParentRouteId
					expectedName := tt.rawData[i].ParentRouteName
					if parentRoute.ParentRouteId != expectedId {
						t.Errorf("ReconstructManyFromRawData()[%d] ParentRouteId = %v, want %v", i, parentRoute.ParentRouteId, expectedId)
					}
					if parentRoute.ParentRouteName.Value() != expectedName {
						t.Errorf("ReconstructManyFromRawData()[%d] ParentRouteName = %v, want %v", i, parentRoute.ParentRouteName.Value(), expectedName)
					}
				}
			}
		})
	}
}
