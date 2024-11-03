package com.tsunacan.expressbustimetableapp.data.model

import kotlinx.serialization.Serializable

@Serializable
data class DepartureTimeAndDestinationApiModel(
    val departureTime: String,
    val destination: String
)
