package com.tsunacan.expressbustimetableapp.data.mapper

import android.util.Log
import com.tsunacan.expressbustimetableapp.data.model.TimeTableApiModel
import com.tsunacan.expressbustimetableapp.models.TimeTable
import com.tsunacan.expressbustimetableapp.models.TimeTableEntry
import java.time.DayOfWeek
import java.time.LocalTime
import java.time.format.DateTimeFormatter
import java.time.format.DateTimeParseException

object TimeTableMapper {
    fun map(timeTableApiModel: TimeTableApiModel): TimeTable {
        return TimeTable(
            parentRouteName = timeTableApiModel.parentRouteName,
            stopName = timeTableApiModel.stopName,
            timeTableEntryList = timeTableApiModel.timeTableEntryList.mapNotNull { timeTableEntryApiModel ->
                departureTimeStringToLocalTime(timeTableEntryApiModel.departureTime)?.let {
                    TimeTableEntry(
                        departureTime = it,
                        destination = timeTableEntryApiModel.destination,
                        dayOfWeekSet = integerSetToDayOfWeekSet(timeTableEntryApiModel.dayOfWeekSet)
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

    private fun integerSetToDayOfWeekSet(dayOfWeekInt: List<Int>): Set<DayOfWeek> {
        return dayOfWeekInt.mapNotNull { dayOfWeekInt ->
            when (dayOfWeekInt) {
                1 -> DayOfWeek.MONDAY
                2 -> DayOfWeek.TUESDAY
                3 -> DayOfWeek.WEDNESDAY
                4 -> DayOfWeek.THURSDAY
                5 -> DayOfWeek.FRIDAY
                6 -> DayOfWeek.SATURDAY
                7 -> DayOfWeek.SUNDAY
                else -> null
            }
        }.toSet()
    }
}
