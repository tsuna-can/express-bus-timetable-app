package com.tsunacan.expressbustimetableapp.data.mapper

import android.util.Log
import com.tsunacan.expressbustimetableapp.data.model.TimeTableApiModel
import com.tsunacan.expressbustimetableapp.data.model.TimeTableEntryApiModel
import com.tsunacan.expressbustimetableapp.models.TimeTable
import com.tsunacan.expressbustimetableapp.models.TimeTableEntry
import java.time.DayOfWeek
import java.time.LocalTime
import java.time.format.DateTimeFormatter
import java.time.format.DateTimeParseException

object TimeTableMapper {
    private val timeFormatter = DateTimeFormatter.ofPattern("HH:mm")

    fun map(timeTableApiModel: TimeTableApiModel): TimeTable = TimeTable(
        parentRouteName = timeTableApiModel.parentRouteName,
        stopName = timeTableApiModel.stopName,
        timeTableEntryList = timeTableApiModel.timeTableEntryList.mapNotNull { mapEntry(it) }
    )

    private fun mapEntry(entry: TimeTableEntryApiModel): TimeTableEntry? {
        val departureTime = departureTimeStringToLocalTime(entry.departureTime) ?: return null
        return TimeTableEntry(
            departureTime = departureTime,
            destination = entry.destination,
            availableDayOfWeek = integerSetToDayOfWeekSet(entry.availableDayOfWeek)
        )
    }

    private fun departureTimeStringToLocalTime(departureTime: String): LocalTime? = try {
        LocalTime.parse(departureTime, timeFormatter)
    } catch (e: DateTimeParseException) {
        Log.i("TimeTableMapper", "Error parsing time: ${e.message}")
        null
    }

    private fun integerSetToDayOfWeekSet(dayOfWeekInts: List<Int>): Set<DayOfWeek> =
        dayOfWeekInts.mapNotNull { it.toDayOfWeek() }.toSet()

    private fun Int.toDayOfWeek(): DayOfWeek? = when (this) {
        1 -> DayOfWeek.MONDAY
        2 -> DayOfWeek.TUESDAY
        3 -> DayOfWeek.WEDNESDAY
        4 -> DayOfWeek.THURSDAY
        5 -> DayOfWeek.FRIDAY
        6 -> DayOfWeek.SATURDAY
        7 -> DayOfWeek.SUNDAY
        else -> null
    }
}
