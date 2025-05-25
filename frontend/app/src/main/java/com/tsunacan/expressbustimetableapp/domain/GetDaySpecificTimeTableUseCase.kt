package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.models.TimeTable
import kotlinx.coroutines.flow.first
import java.time.LocalDate
import javax.inject.Inject

class GetDaySpecificTimeTableUseCase @Inject constructor(
    private val timeTableRepository: TimeTableRepository
) {
    /**
     * Get the time table for the selected bus stop, selected day of the week and selected route
     *
     * @param parentRouteId The id of the parent route
     * @param busStopId The id of the bus stop
     * @param today The selected day of the week
     * @return The time table for the selected bus stop, selected day of the week and selected route
     */
    suspend operator fun invoke(
        parentRouteId: String,
        busStopId: String,
        today: LocalDate
    ): TimeTable {
        val busStopTimeTable = timeTableRepository.getTimeTable(parentRouteId, busStopId).first()

        val dayOfWeek = today.dayOfWeek

        // Filter the time table based on the selected day of the week
        val filteredTimeTable = busStopTimeTable.timeTableEntryList.filter {
            it.availableDayOfWeek.contains(dayOfWeek)
        }

        // Sort the time table based on the departure time
        val sortedTimeTable = filteredTimeTable.sortedBy { it.departureTime }

        return busStopTimeTable.copy(timeTableEntryList = sortedTimeTable)
    }
}
