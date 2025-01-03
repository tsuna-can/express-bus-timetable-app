package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.models.TimeTable
import javax.inject.Inject
import java.time.LocalDateTime

class GetUpcomingTimeTableUseCase @Inject constructor(
    private val getDaySpecificTimeTableUseCase: GetDaySpecificTimeTableUseCase
) {
    suspend operator fun invoke(
        parentRouteId: String,
        busStopId: String
    ): TimeTable {
        val busStopTimeTable = getDaySpecificTimeTableUseCase.invoke(parentRouteId, busStopId)

        val currentTime = LocalDateTime.now().toLocalTime()

        // Get the index of the first departure time that is greater than the current time
        val closestDepartureTimeIndex =
            busStopTimeTable.departureTimeAndDestinationList.indexOfFirst {
                it.departureTime > currentTime
            }

        // In current implementation, if there is no upcoming departure time, return empty list
        // TODO return the next day's first departure time
        if (closestDepartureTimeIndex == -1) {
            return busStopTimeTable.copy(departureTimeAndDestinationList = emptyList())
        }

        // Extract 3 upcoming departure time and destination
        // If the closest departure time is the last one, we will get the last 3 departure time
        val upComingTimeTable = busStopTimeTable.departureTimeAndDestinationList.subList(
            closestDepartureTimeIndex,
            minOf(
                busStopTimeTable.departureTimeAndDestinationList.size,
                closestDepartureTimeIndex + 3
            )
        )

        return busStopTimeTable.copy(departureTimeAndDestinationList = upComingTimeTable)
    }
}
