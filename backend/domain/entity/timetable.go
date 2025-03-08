package entity

type Timetable struct {
	ParentRouteId    string
	BusStopId        string
	TimetableEntries []TimetableEntry
}

func NewTimetable(
	parentRouteId string,
	busStopId string,
	timetableEntries []TimetableEntry) *Timetable {
	return &Timetable{
		ParentRouteId:    parentRouteId,
		BusStopId:        busStopId,
		TimetableEntries: timetableEntries,
	}
}
