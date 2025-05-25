package com.tsunacan.expressbustimetableapp.models

data class BusStop(
    val parentRouteId: String,
    val parentRouteName: String,
    val stopId: String,
    val stopName: String,
)
