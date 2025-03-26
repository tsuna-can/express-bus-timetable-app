package response

import (
  "github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type TimetableResponse struct {
  ParentRouteId string `json:"parent_route_id"`
  BusStopId     string `json:"bus_stop_id"`
  Timetables    []TimetableEntry`json:"timetable_entry"`
}

type TimetableEntry struct {
  DepartureTime   string   `json:"departure_time"`
  DestinationName string   `json:"destination_name"`
  OperationDays   []int `json:"operation_days"`
}

func NewTimetableResponse(timetable entity.Timetable) *TimetableResponse {
  timetableEntries := make([]TimetableEntry, 0, len(timetable.TimetableEntries))
  for _, te := range timetable.TimetableEntries {
    timetableEntries = append(timetableEntries, TimetableEntry{
      DepartureTime:   te.DepartureTime.Value(),
      DestinationName: te.DestinationName.Value(),
      OperationDays:   te.OperationDaysAsIntSlice(),
    })
  }
  return &TimetableResponse{
    ParentRouteId: timetable.ParentRouteId,
    BusStopId:     timetable.BusStopId,
    Timetables:    timetableEntries,
  }
}

