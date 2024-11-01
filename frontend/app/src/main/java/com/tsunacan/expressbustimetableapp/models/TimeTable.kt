package com.tsunacan.expressbustimetableapp.models

/**
 * Time table for a bus stop
 */
data class TimeTable (
    val parentRouteName: String = "",
    val stopName: String = "",
    val departureTimeAndDestinationList: List<DepartureTimeAndDestination> = emptyList()
)
