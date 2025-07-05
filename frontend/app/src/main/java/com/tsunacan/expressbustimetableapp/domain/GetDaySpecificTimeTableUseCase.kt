package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.data.repository.TimetableRepository
import com.tsunacan.expressbustimetableapp.models.Timetable
import kotlinx.coroutines.flow.first
import java.time.LocalDate
import javax.inject.Inject

class GetDaySpecificTimetableUseCase @Inject constructor(
    private val timetableRepository: TimetableRepository
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
    ): Timetable {
        val busStopTimetable = timetableRepository.getTimetable(parentRouteId, busStopId).first()

        val dayOfWeek = today.dayOfWeek

        // Filter the time table based on the selected day of the week
        val filteredTimetable = busStopTimetable.timetableEntryList.filter {
            it.availableDayOfWeek.contains(dayOfWeek)
        }

        // Sort the time table based on the departure time
        val sortedTimetable = filteredTimetable.sortedBy { it.departureTime }

        return busStopTimetable.copy(timetableEntryList = sortedTimetable)
    }
}
