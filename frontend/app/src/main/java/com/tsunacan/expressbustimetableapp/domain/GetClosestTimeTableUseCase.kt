package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination
import kotlinx.coroutines.flow.first
import javax.inject.Inject
import java.time.LocalDateTime

class GetClosestTimeTableUseCase @Inject constructor(
    private val timeTableRepository: TimeTableRepository
) {

    suspend operator fun invoke(
        parentRouteId: String,
        busStopId: String
    ): List<DepartureTimeAndDestination> {
        val timeTable = timeTableRepository.getTimeTable(parentRouteId, busStopId).first()

        val currentTime = LocalDateTime.now().toLocalTime()

        // Get the index of the first departure time that is greater than the current time
        val closestDepartureTimeIndex = timeTable.indexOfFirst {
            it.departureTime > currentTime
        }

        // Return 3 departure times after the closest departure time
        return timeTable.subList(
            closestDepartureTimeIndex,
            closestDepartureTimeIndex + 3
        )
    }
}
