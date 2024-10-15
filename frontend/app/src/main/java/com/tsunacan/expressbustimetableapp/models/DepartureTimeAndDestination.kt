package com.tsunacan.expressbustimetableapp.models

import java.time.LocalTime

data class DepartureTimeAndDestination(
    val departureTime: LocalTime,
    val destination: String
)
