package com.tsunacan.expressbustimetableapp.data.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class TimetableEntryApiModel(
    @SerialName("departure_time")
    val departureTime: String,
    @SerialName("destination_name")
    val destinationName: String,
    @SerialName("operation_days")
    val operationDays: List<Int>
)
