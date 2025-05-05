package entity

import "github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"

type Timetable struct {
	ParentRouteId    string
	ParentRouteName  vo.ParentRouteName
	BusStopId        string
	BusStopName      vo.BusStopName
	TimetableEntries []TimetableEntry
}

func NewTimetable(
	parentRouteId string,
	parentRouteName vo.ParentRouteName,
	busStopId string,
	busStopName vo.BusStopName,
	timetableEntries []TimetableEntry) *Timetable {
	return &Timetable{
		ParentRouteId:    parentRouteId,
		ParentRouteName:  parentRouteName,
		BusStopId:        busStopId,
		BusStopName:      busStopName,
		TimetableEntries: timetableEntries,
	}
}
