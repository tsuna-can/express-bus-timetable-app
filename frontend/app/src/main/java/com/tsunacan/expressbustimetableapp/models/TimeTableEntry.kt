package com.tsunacan.expressbustimetableapp.models

import java.time.DayOfWeek
import java.time.LocalTime

data class TimeTableEntry(
    val departureTime: LocalTime,
    val destination: String,
    val dayOfWeekSet: Set<DayOfWeek>
)
