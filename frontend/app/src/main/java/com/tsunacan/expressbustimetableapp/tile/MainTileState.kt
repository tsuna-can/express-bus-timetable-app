package com.tsunacan.expressbustimetableapp.tile

import com.tsunacan.expressbustimetableapp.models.TimetableEntry

data class MainTileState(
    val parentRouteId: String = "",
    val parentRouteName: String = "",
    val stopId: String = "",
    val stopName: String = "",
    val timetableEntryList: List<TimetableEntry> = emptyList()
)
