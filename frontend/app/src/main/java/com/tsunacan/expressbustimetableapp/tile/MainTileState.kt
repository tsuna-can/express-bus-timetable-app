package com.tsunacan.expressbustimetableapp.tile

import com.tsunacan.expressbustimetableapp.models.TimeTableEntry

data class MainTileState(
    val parentRouteName: String = "",
    val stopName: String = "",
    val timeTableEntryList: List<TimeTableEntry> = emptyList()
)
