package com.tsunacan.expressbustimetableapp.data.model

import kotlinx.serialization.Serializable

@Serializable
data class TimeTableApiModel (
    val parentRouteId: String,
    val parentRouteName: String,
    val stopId: String,
    val stopName: String,
    val departureTimeAndDestinationList: List<DepartureTimeAndDestinationApiModel>
)

