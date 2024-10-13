package com.tsunacan.expressbustimetableapp.tile

import com.tsunacan.expressbustimetableapp.models.Trip

data class MainTileState (
    val stopName: String,
    val timeTable: List<Trip>
)
