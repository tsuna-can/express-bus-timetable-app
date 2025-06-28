package response

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type TimetableResponse struct {
	ParentRouteId   string           `json:"parent_route_id"`
	ParentRouteName string           `json:"parent_route_name"`
	BusStopId       string           `json:"bus_stop_id"`
	BusStopName     string           `json:"bus_stop_name"`
	Timetables      []TimetableEntry `json:"timetable_entry"`
}

type TimetableEntry struct {
	DepartureTime   string `json:"departure_time"`
	DestinationName string `json:"destination_name"`
	OperationDays   []int  `json:"operation_days"`
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
		ParentRouteId:   timetable.ParentRouteId,
		ParentRouteName: timetable.ParentRouteName.Value(),
		BusStopId:       timetable.BusStopId,
		BusStopName:     timetable.BusStopName.Value(),
		Timetables:      timetableEntries,
	}
}
