package com.tsunacan.expressbustimetableapp.models

/**
 * Time table for a bus stop
 */
data class TimeTable(
    val parentRouteId: String = "",
    val parentRouteName: String = "",
    val stopId: String = "",
    val stopName: String = "",
    val departureTimeAndDestinationList: List<DepartureTimeAndDestination> = emptyList()
)
