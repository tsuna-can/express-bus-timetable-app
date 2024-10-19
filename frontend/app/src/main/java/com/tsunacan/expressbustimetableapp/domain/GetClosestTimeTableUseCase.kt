package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination
import kotlinx.coroutines.flow.first
import javax.inject.Inject

class GetClosestTimeTableUseCase @Inject constructor(
    private val timeTableRepository: TimeTableRepository
) {

    suspend operator fun invoke(
    ): List<DepartureTimeAndDestination> {
        return timeTableRepository.getTimeTable("test", "test").first()
    }
}
