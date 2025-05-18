package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.models.TimeTable
import java.time.LocalDate
import javax.inject.Inject
import java.time.LocalDateTime

class GetUpcomingTimeTableUseCase @Inject constructor(
    private val getDaySpecificTimeTableUseCase: GetDaySpecificTimeTableUseCase
) {
    suspend operator fun invoke(
        parentRouteId: String,
        busStopId: String,
        today: LocalDate
    ): TimeTable {
        val busStopTimeTable =
            getDaySpecificTimeTableUseCase.invoke(parentRouteId, busStopId, today)

        val currentTime = LocalDateTime.now().toLocalTime()

        // Get the index of the first departure time that is greater than the current time
        val closestDepartureTimeIndex =
            busStopTimeTable.timeTableEntryList.indexOfFirst {
                it.departureTime >= currentTime
            }

        // In current implementation, if there is no upcoming departure time, return empty list
        // TODO return the next day's first departure time
        if (closestDepartureTimeIndex == -1) {
            return busStopTimeTable.copy(timeTableEntryList = emptyList())
        }

        // Extract 3 upcoming departure time and destination
        // If the closest departure time is the last one, we will get the last 3 departure time
        val upComingTimeTable = busStopTimeTable.timeTableEntryList.subList(
            closestDepartureTimeIndex,
            minOf(
                busStopTimeTable.timeTableEntryList.size,
                closestDepartureTimeIndex + 3
            )
        )

        return busStopTimeTable.copy(timeTableEntryList = upComingTimeTable)
    }
}
