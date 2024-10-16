package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flowOf
import java.time.LocalDateTime

class TimeTableRepository {

    suspend fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): Flow<List<DepartureTimeAndDestination>> {
        // As of now, we are returning a dummy list of DepartureTimeAndDestination
        return flowOf(
            listOf(
                DepartureTimeAndDestination(
                    LocalDateTime.now().plusMinutes(10).toLocalTime(),
                    "Shibuya"
                ),
                DepartureTimeAndDestination(
                    LocalDateTime.now().plusMinutes(20).toLocalTime(),
                    "Shibuya"
                ),
                DepartureTimeAndDestination(
                    LocalDateTime.now().plusMinutes(30).toLocalTime(),
                    "Shibuya"
                ),
            )
        )
    }

}