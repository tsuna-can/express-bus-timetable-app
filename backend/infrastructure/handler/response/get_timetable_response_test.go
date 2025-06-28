package response

import (
	"testing"
	"time"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

func TestNewTimetableResponse(t *testing.T) {
	// Given
	parentRouteName, _ := vo.NewParentRouteName("テスト路線")
	busStopName, _ := vo.NewBusStopName("テストバス停")
	departureTime, _ := vo.NewDepartureTime("08:30")
	destinationName, _ := vo.NewDestinationName("テスト行き先")

	operationDays := map[vo.OperationDay]struct{}{
		*vo.NewOperationDay(time.Monday):    {},
		*vo.NewOperationDay(time.Tuesday):   {},
		*vo.NewOperationDay(time.Wednesday): {},
	}

	timetableEntry := entity.NewTimetableEntry(
		*departureTime,
		operationDays,
		*destinationName,
	)

	timetable := entity.NewTimetable(
		"route-001",
		*parentRouteName,
		"stop-001",
		*busStopName,
		[]entity.TimetableEntry{*timetableEntry},
	)

	// When
	response := NewTimetableResponse(*timetable)

	// Then
	if response.ParentRouteId != "route-001" {
		t.Errorf("Expected ParentRouteId to be 'route-001', got '%s'", response.ParentRouteId)
	}

	if response.ParentRouteName != "テスト路線" {
		t.Errorf("Expected ParentRouteName to be 'テスト路線', got '%s'", response.ParentRouteName)
	}

	if response.BusStopId != "stop-001" {
		t.Errorf("Expected BusStopId to be 'stop-001', got '%s'", response.BusStopId)
	}

	if response.BusStopName != "テストバス停" {
		t.Errorf("Expected BusStopName to be 'テストバス停', got '%s'", response.BusStopName)
	}

	if len(response.Timetables) != 1 {
		t.Errorf("Expected 1 timetable entry, got %d", len(response.Timetables))
	}

	timetableResp := response.Timetables[0]
	if timetableResp.DepartureTime != "08:30" {
		t.Errorf("Expected DepartureTime to be '08:30', got '%s'", timetableResp.DepartureTime)
	}

	if timetableResp.DestinationName != "テスト行き先" {
		t.Errorf("Expected DestinationName to be 'テスト行き先', got '%s'", timetableResp.DestinationName)
	}

	// OperationDays should contain 1, 2, 3 (Monday, Tuesday, Wednesday)
	if len(timetableResp.OperationDays) != 3 {
		t.Errorf("Expected 3 operation days, got %d", len(timetableResp.OperationDays))
	}

	expectedDays := map[int]bool{1: true, 2: true, 3: true}
	for _, day := range timetableResp.OperationDays {
		if !expectedDays[day] {
			t.Errorf("Unexpected operation day: %d", day)
		}
		delete(expectedDays, day)
	}
	if len(expectedDays) > 0 {
		t.Errorf("Missing operation days: %v", expectedDays)
	}
}

func TestNewTimetableResponseEmptyTimetableEntries(t *testing.T) {
	// Given
	parentRouteName, _ := vo.NewParentRouteName("テスト路線")
	busStopName, _ := vo.NewBusStopName("テストバス停")

	timetable := entity.NewTimetable(
		"route-001",
		*parentRouteName,
		"stop-001",
		*busStopName,
		[]entity.TimetableEntry{},
	)

	// When
	response := NewTimetableResponse(*timetable)

	// Then
	if len(response.Timetables) != 0 {
		t.Errorf("Expected 0 timetable entries, got %d", len(response.Timetables))
	}
}

func TestNewTimetableResponseMultipleTimetableEntries(t *testing.T) {
	// Given
	parentRouteName, _ := vo.NewParentRouteName("テスト路線")
	busStopName, _ := vo.NewBusStopName("テストバス停")

	departureTime1, _ := vo.NewDepartureTime("08:30")
	destinationName1, _ := vo.NewDestinationName("テスト行き先1")
	operationDays1 := map[vo.OperationDay]struct{}{
		*vo.NewOperationDay(time.Monday): {},
	}

	departureTime2, _ := vo.NewDepartureTime("09:30")
	destinationName2, _ := vo.NewDestinationName("テスト行き先2")
	operationDays2 := map[vo.OperationDay]struct{}{
		*vo.NewOperationDay(time.Sunday): {},
	}

	timetableEntry1 := entity.NewTimetableEntry(*departureTime1, operationDays1, *destinationName1)
	timetableEntry2 := entity.NewTimetableEntry(*departureTime2, operationDays2, *destinationName2)

	timetable := entity.NewTimetable(
		"route-001",
		*parentRouteName,
		"stop-001",
		*busStopName,
		[]entity.TimetableEntry{*timetableEntry1, *timetableEntry2},
	)

	// When
	response := NewTimetableResponse(*timetable)

	// Then
	if len(response.Timetables) != 2 {
		t.Errorf("Expected 2 timetable entries, got %d", len(response.Timetables))
	}

	// 順序は保証されないので、両方の要素が含まれていることを確認
	foundEntry1 := false
	foundEntry2 := false

	for _, entry := range response.Timetables {
		if entry.DepartureTime == "08:30" && entry.DestinationName == "テスト行き先1" {
			foundEntry1 = true
			if len(entry.OperationDays) != 1 || entry.OperationDays[0] != 1 {
				t.Errorf("Entry1: Expected operation day [1], got %v", entry.OperationDays)
			}
		} else if entry.DepartureTime == "09:30" && entry.DestinationName == "テスト行き先2" {
			foundEntry2 = true
			if len(entry.OperationDays) != 1 || entry.OperationDays[0] != 7 {
				t.Errorf("Entry2: Expected operation day [7], got %v", entry.OperationDays)
			}
		}
	}

	if !foundEntry1 {
		t.Error("Entry1 not found in response")
	}
	if !foundEntry2 {
		t.Error("Entry2 not found in response")
	}
}
