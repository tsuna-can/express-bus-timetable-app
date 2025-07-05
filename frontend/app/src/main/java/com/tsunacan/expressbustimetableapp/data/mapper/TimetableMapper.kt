package com.tsunacan.expressbustimetableapp.data.mapper

import android.util.Log
import com.tsunacan.expressbustimetableapp.data.model.TimetableApiModel
import com.tsunacan.expressbustimetableapp.data.model.TimetableEntryApiModel
import com.tsunacan.expressbustimetableapp.models.Timetable
import com.tsunacan.expressbustimetableapp.models.TimetableEntry
import java.time.DayOfWeek
import java.time.LocalTime
import java.time.format.DateTimeFormatter
import java.time.format.DateTimeParseException

object TimetableMapper {
    private val timeFormatter = DateTimeFormatter.ofPattern("HH:mm")

    fun mapToTimetable(timetableApiModel: TimetableApiModel): Timetable = Timetable(
        parentRouteId = timetableApiModel.parentRouteId,
        parentRouteName = timetableApiModel.parentRouteName,
        stopId = timetableApiModel.busStopId,
        stopName = timetableApiModel.busStopName,
        timetableEntryList = timetableApiModel.timetableEntry.mapNotNull { mapEntry(it) }
    )

    private fun mapEntry(entry: TimetableEntryApiModel): TimetableEntry? {
        val departureTime = departureTimeStringToLocalTime(entry.departureTime) ?: return null
        return TimetableEntry(
            departureTime = departureTime,
            destination = entry.destinationName,
            availableDayOfWeek = integerSetToDayOfWeekSet(entry.operationDays)
        )
    }

    private fun departureTimeStringToLocalTime(departureTime: String): LocalTime? = try {
        LocalTime.parse(departureTime, timeFormatter)
    } catch (e: DateTimeParseException) {
        Log.i("TimetableMapper", "Error parsing time: ${e.message}")
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
