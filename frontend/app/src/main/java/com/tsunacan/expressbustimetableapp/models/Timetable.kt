package com.tsunacan.expressbustimetableapp.models

/**
 * Time table for a bus stop
 */
data class Timetable(
    val parentRouteId: String = "",
    val parentRouteName: String = "",
    val stopId: String = "",
    val stopName: String = "",
    val timetableEntryList: List<TimetableEntry> = emptyList()
)
