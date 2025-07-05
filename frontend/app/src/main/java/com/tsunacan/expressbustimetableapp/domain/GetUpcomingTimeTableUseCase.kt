package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.models.Timetable
import java.time.LocalDate
import java.time.LocalDateTime
import javax.inject.Inject

class GetUpcomingTimetableUseCase @Inject constructor(
    private val getDaySpecificTimetableUseCase: GetDaySpecificTimetableUseCase
) {
    suspend operator fun invoke(
        parentRouteId: String,
        busStopId: String,
        today: LocalDate
    ): Timetable {
        val busStopTimetable =
            getDaySpecificTimetableUseCase.invoke(parentRouteId, busStopId, today)

        val currentTime = LocalDateTime.now().toLocalTime()

        // Get the index of the first departure time that is greater than the current time
        val closestDepartureTimeIndex =
            busStopTimetable.timetableEntryList.indexOfFirst {
                it.departureTime >= currentTime
            }

        // In current implementation, if there is no upcoming departure time, return empty list
        // TODO return the next day's first departure time
        if (closestDepartureTimeIndex == -1) {
            return busStopTimetable.copy(timetableEntryList = emptyList())
        }

        // Extract 3 upcoming departure time and destination
        // If the closest departure time is the last one, we will get the last 3 departure time
        val upComingTimetable = busStopTimetable.timetableEntryList.subList(
            closestDepartureTimeIndex,
            minOf(
                busStopTimetable.timetableEntryList.size,
                closestDepartureTimeIndex + 3
            )
        )

        return busStopTimetable.copy(timetableEntryList = upComingTimetable)
    }
}
