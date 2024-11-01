package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.models.TimeTable
import kotlinx.coroutines.flow.first
import javax.inject.Inject
import java.time.LocalDateTime

class GetUpcomingTimeTableUseCase @Inject constructor(
    private val timeTableRepository: TimeTableRepository
) {

    suspend operator fun invoke(
        parentRouteId: String,
        busStopId: String
    ): TimeTable {
        val busStopTimeTable = timeTableRepository.getTimeTable(parentRouteId, busStopId).first()

        val currentTime = LocalDateTime.now().toLocalTime()

        // Get the index of the first departure time that is greater than the current time
        val closestDepartureTimeIndex =
            busStopTimeTable.departureTimeAndDestinationList.indexOfFirst {
                it.departureTime > currentTime
            }

        // Extract 3 upcoming departure time and destination
        val upComingTimeTable = busStopTimeTable.departureTimeAndDestinationList.subList(
            closestDepartureTimeIndex,
            closestDepartureTimeIndex + 3
        )

        return busStopTimeTable.copy(departureTimeAndDestinationList = upComingTimeTable)
    }
}
