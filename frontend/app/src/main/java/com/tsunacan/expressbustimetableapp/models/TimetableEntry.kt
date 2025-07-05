package com.tsunacan.expressbustimetableapp.models

import java.time.DayOfWeek
import java.time.LocalTime

data class TimetableEntry(
    val departureTime: LocalTime,
    val destination: String,
    val availableDayOfWeek: Set<DayOfWeek>
)
