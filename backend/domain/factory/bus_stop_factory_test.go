package factory

import (
	"testing"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewBusStopFactory(t *testing.T) {
	factory := NewBusStopFactory()
	if factory == nil {
		t.Error("NewBusStopFactory() should not return nil")
	}
}

func TestBusStopFactory_ReconstructFromRawData(t *testing.T) {
	factory := NewBusStopFactory()

	tests := []struct {
		name    string
		rawData BusStopRawData
		want    *entity.BusStop
		wantErr bool
	}{
		{
			name: "正常系：有効なデータ",
			rawData: BusStopRawData{
				BusStopId:   "stop001",
				BusStopName: "新宿駅",
			},
			want: func() *entity.BusStop {
				name, _ := vo.NewBusStopName("新宿駅")
				return entity.NewBusStop("stop001", *name)
			}(),
			wantErr: false,
		},
		{
			name: "異常系：空のバス停名",
			rawData: BusStopRawData{
				BusStopId:   "stop001",
				BusStopName: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系：空白のみのバス停名",
			rawData: BusStopRawData{
				BusStopId:   "stop001",
				BusStopName: "   ",
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
				if got.BusStopId != tt.want.BusStopId {
					t.Errorf("ReconstructFromRawData() BusStopId = %v, want %v", got.BusStopId, tt.want.BusStopId)
				}
				if got.BusStopName.Value() != tt.want.BusStopName.Value() {
					t.Errorf("ReconstructFromRawData() BusStopName = %v, want %v", got.BusStopName.Value(), tt.want.BusStopName.Value())
				}
			}
		})
	}
}

func TestBusStopFactory_ReconstructManyFromRawData(t *testing.T) {
	factory := NewBusStopFactory()

	tests := []struct {
		name     string
		rawData  []BusStopRawData
		wantLen  int
		wantErr  bool
		errIndex int // エラーが発生する予想インデックス
	}{
		{
			name: "正常系：複数の有効なデータ",
			rawData: []BusStopRawData{
				{BusStopId: "stop001", BusStopName: "新宿駅"},
				{BusStopId: "stop002", BusStopName: "渋谷駅"},
				{BusStopId: "stop003", BusStopName: "池袋駅"},
			},
			wantLen: 3,
			wantErr: false,
		},
		{
			name:    "正常系：空のスライス",
			rawData: []BusStopRawData{},
			wantLen: 0,
			wantErr: false,
		},
		{
			name: "異常系：2番目のデータが無効",
			rawData: []BusStopRawData{
				{BusStopId: "stop001", BusStopName: "新宿駅"},
				{BusStopId: "stop002", BusStopName: ""}, // 無効
				{BusStopId: "stop003", BusStopName: "池袋駅"},
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
				for i, busStop := range got {
					expectedId := tt.rawData[i].BusStopId
					expectedName := tt.rawData[i].BusStopName
					if busStop.BusStopId != expectedId {
						t.Errorf("ReconstructManyFromRawData()[%d] BusStopId = %v, want %v", i, busStop.BusStopId, expectedId)
					}
					if busStop.BusStopName.Value() != expectedName {
						t.Errorf("ReconstructManyFromRawData()[%d] BusStopName = %v, want %v", i, busStop.BusStopName.Value(), expectedName)
					}
				}
			}
		})
	}
}
