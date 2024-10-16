package com.tsunacan.expressbustimetableapp.tile

import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination

data class MainTileState (
    val parentRouteName: String = "",
    val stopName: String = "",
    val timeTable: List<DepartureTimeAndDestination> = emptyList()
)
