package com.tsunacan.expressbustimetableapp.data.model

import kotlinx.serialization.Serializable

@Serializable
data class TimeTableEntryApiModel( // TODO rename this
    val departureTime: String,
    val destination: String,
    val dayOfWeekSet: List<Int>
)
