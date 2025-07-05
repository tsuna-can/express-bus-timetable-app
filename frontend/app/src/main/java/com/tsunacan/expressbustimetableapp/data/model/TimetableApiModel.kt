package com.tsunacan.expressbustimetableapp.data.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class TimetableApiModel(
    @SerialName("parent_route_id")
    val parentRouteId: String,
    @SerialName("parent_route_name")
    val parentRouteName: String,
    @SerialName("bus_stop_id")
    val busStopId: String,
    @SerialName("bus_stop_name")
    val busStopName: String,
    @SerialName("timetable_entry")
    val timetableEntry: List<TimetableEntryApiModel>
)
