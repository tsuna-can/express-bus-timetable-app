package com.tsunacan.expressbustimetableapp.data.mapper

import android.util.Log
import com.tsunacan.expressbustimetableapp.data.model.TimeTableApiModel
import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination
import com.tsunacan.expressbustimetableapp.models.TimeTable
import java.time.LocalTime
import java.time.format.DateTimeFormatter
import java.time.format.DateTimeParseException

object TimeTableMapper {
    fun map(timeTableApiModel: TimeTableApiModel): TimeTable {
        return TimeTable(
            parentRouteName = timeTableApiModel.parentRouteName,
            stopName = timeTableApiModel.stopName,
            departureTimeAndDestinationList = timeTableApiModel.departureTimeAndDestinationList.mapNotNull { departureTimeAndDestinationApiModel ->
                departureTimeStringToLocalTime(departureTimeAndDestinationApiModel.departureTime)?.let {
                    DepartureTimeAndDestination(
                        departureTime = it,
                        destination = departureTimeAndDestinationApiModel.destination
                    )
                }
            }
        )
    }

    private fun departureTimeStringToLocalTime(departureTime: String): LocalTime? {
        val formatter = DateTimeFormatter.ofPattern("HH:mm")

        return try {
            LocalTime.parse(departureTime, formatter)
        } catch (e: DateTimeParseException) {
            Log.i("TimeTableMapper", "Error parsing time: ${e.message}")
            return null
        }
    }
}
