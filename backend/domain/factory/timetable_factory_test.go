package factory

import (
	"testing"
)

func TestNewTimetableFactory(t *testing.T) {
	factory := NewTimetableFactory()
	if factory == nil {
		t.Error("NewTimetableFactory() should not return nil")
	}
}

func TestTimetableFactory_ReconstructFromRawData(t *testing.T) {
	factory := NewTimetableFactory()

	tests := []struct {
		name    string
		rawData TimetableRawData
		wantErr bool
	}{
		{
			name: "正常系：有効なデータ",
			rawData: TimetableRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "山手線",
				BusStopId:       "stop001",
				BusStopName:     "新宿駅",
				Entries: []TimetableEntryRawData{
					{
						DepartureTime:   "08:30",
						DestinationName: "品川方面",
						Monday:          true,
						Tuesday:         true,
						Wednesday:       true,
						Thursday:        true,
						Friday:          true,
						Saturday:        false,
						Sunday:          false,
					},
					{
						DepartureTime:   "09:00",
						DestinationName: "品川方面",
						Monday:          false,
						Tuesday:         false,
						Wednesday:       false,
						Thursday:        false,
						Friday:          false,
						Saturday:        true,
						Sunday:          true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "正常系：エントリなし",
			rawData: TimetableRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "山手線",
				BusStopId:       "stop001",
				BusStopName:     "新宿駅",
				Entries:         []TimetableEntryRawData{},
			},
			wantErr: false,
		},
		{
			name: "異常系：空の路線名",
			rawData: TimetableRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "",
				BusStopId:       "stop001",
				BusStopName:     "新宿駅",
				Entries:         []TimetableEntryRawData{},
			},
			wantErr: true,
		},
		{
			name: "異常系：空のバス停名",
			rawData: TimetableRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "山手線",
				BusStopId:       "stop001",
				BusStopName:     "",
				Entries:         []TimetableEntryRawData{},
			},
			wantErr: true,
		},
		{
			name: "異常系：無効な出発時刻",
			rawData: TimetableRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "山手線",
				BusStopId:       "stop001",
				BusStopName:     "新宿駅",
				Entries: []TimetableEntryRawData{
					{
						DepartureTime:   "25:00", // 無効な時刻
						DestinationName: "品川方面",
						Monday:          true,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "異常系：空の目的地名",
			rawData: TimetableRawData{
				ParentRouteId:   "route001",
				ParentRouteName: "山手線",
				BusStopId:       "stop001",
				BusStopName:     "新宿駅",
				Entries: []TimetableEntryRawData{
					{
						DepartureTime:   "08:30",
						DestinationName: "", // 空の目的地名
						Monday:          true,
					},
				},
			},
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
				// 基本フィールドの確認
				if got.ParentRouteId != tt.rawData.ParentRouteId {
					t.Errorf("ReconstructFromRawData() ParentRouteId = %v, want %v", got.ParentRouteId, tt.rawData.ParentRouteId)
				}
				if got.ParentRouteName.Value() != tt.rawData.ParentRouteName {
					t.Errorf("ReconstructFromRawData() ParentRouteName = %v, want %v", got.ParentRouteName.Value(), tt.rawData.ParentRouteName)
				}
				if got.BusStopId != tt.rawData.BusStopId {
					t.Errorf("ReconstructFromRawData() BusStopId = %v, want %v", got.BusStopId, tt.rawData.BusStopId)
				}
				if got.BusStopName.Value() != tt.rawData.BusStopName {
					t.Errorf("ReconstructFromRawData() BusStopName = %v, want %v", got.BusStopName.Value(), tt.rawData.BusStopName)
				}
				// エントリ数の確認
				if len(got.TimetableEntries) != len(tt.rawData.Entries) {
					t.Errorf("ReconstructFromRawData() TimetableEntries length = %v, want %v", len(got.TimetableEntries), len(tt.rawData.Entries))
				}
			}
		})
	}
}

func TestTimetableFactory_createOperationDays(t *testing.T) {
	factory := &timetableFactory{}

	tests := []struct {
		name    string
		rawData TimetableEntryRawData
		want    int // 期待される OperationDays の数
	}{
		{
			name: "平日のみ",
			rawData: TimetableEntryRawData{
				Monday:    true,
				Tuesday:   true,
				Wednesday: true,
				Thursday:  true,
				Friday:    true,
				Saturday:  false,
				Sunday:    false,
			},
			want: 5,
		},
		{
			name: "週末のみ",
			rawData: TimetableEntryRawData{
				Monday:    false,
				Tuesday:   false,
				Wednesday: false,
				Thursday:  false,
				Friday:    false,
				Saturday:  true,
				Sunday:    true,
			},
			want: 2,
		},
		{
			name: "全曜日",
			rawData: TimetableEntryRawData{
				Monday:    true,
				Tuesday:   true,
				Wednesday: true,
				Thursday:  true,
				Friday:    true,
				Saturday:  true,
				Sunday:    true,
			},
			want: 7,
		},
		{
			name: "曜日なし",
			rawData: TimetableEntryRawData{
				Monday:    false,
				Tuesday:   false,
				Wednesday: false,
				Thursday:  false,
				Friday:    false,
				Saturday:  false,
				Sunday:    false,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := factory.createOperationDays(tt.rawData)
			if len(got) != tt.want {
				t.Errorf("createOperationDays() length = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestTimetableFactory_createTimetableEntry(t *testing.T) {
	factory := &timetableFactory{}

	tests := []struct {
		name    string
		rawData TimetableEntryRawData
		wantErr bool
	}{
		{
			name: "正常系：有効なデータ",
			rawData: TimetableEntryRawData{
				DepartureTime:   "08:30",
				DestinationName: "品川方面",
				Monday:          true,
				Friday:          true,
			},
			wantErr: false,
		},
		{
			name: "異常系：無効な出発時刻",
			rawData: TimetableEntryRawData{
				DepartureTime:   "25:00",
				DestinationName: "品川方面",
				Monday:          true,
			},
			wantErr: true,
		},
		{
			name: "異常系：空の目的地名",
			rawData: TimetableEntryRawData{
				DepartureTime:   "08:30",
				DestinationName: "",
				Monday:          true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := factory.createTimetableEntry(tt.rawData)
			if (err != nil) != tt.wantErr {
				t.Errorf("createTimetableEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Error("createTimetableEntry() should not return nil when no error expected")
					return
				}
				if got.DepartureTime.Value() != tt.rawData.DepartureTime {
					t.Errorf("createTimetableEntry() DepartureTime = %v, want %v", got.DepartureTime.Value(), tt.rawData.DepartureTime)
				}
				if got.DestinationName.Value() != tt.rawData.DestinationName {
					t.Errorf("createTimetableEntry() DestinationName = %v, want %v", got.DestinationName.Value(), tt.rawData.DestinationName)
				}
			}
		})
	}
}
