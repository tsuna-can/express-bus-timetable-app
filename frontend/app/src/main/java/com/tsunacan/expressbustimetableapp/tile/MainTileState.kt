package com.tsunacan.expressbustimetableapp.tile

import com.tsunacan.expressbustimetableapp.models.TimeTableEntry

data class MainTileState(
    val parentRouteId : String = "",
    val parentRouteName: String = "",
    val stopId: String = "",
    val stopName: String = "",
    val timeTableEntryList: List<TimeTableEntry> = emptyList()
)
