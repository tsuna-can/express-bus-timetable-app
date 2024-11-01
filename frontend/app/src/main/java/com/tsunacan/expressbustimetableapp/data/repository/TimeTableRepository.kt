package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.models.TimeTable
import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flowOf
import java.time.LocalDateTime

interface TimeTableRepository {
    suspend fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): Flow<TimeTable>
}

class TimeTableRepositoryImpl : TimeTableRepository {

    override suspend fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): Flow<TimeTable> {
        // As of now, we are returning a dummy list of DepartureTimeAndDestination
        return flowOf(
            TimeTable(
                parentRouteName = "Route 1",
                stopName = "Bus Stop 1",
                departureTimeAndDestinationList =
                listOf(
                    DepartureTimeAndDestination(
                        LocalDateTime.now().minusMinutes(20).toLocalTime(),
                        "Shibuya"
                    ),
                    DepartureTimeAndDestination(
                        LocalDateTime.now().minusMinutes(10).toLocalTime(),
                        "Shibuya"
                    ),
                    DepartureTimeAndDestination(
                        LocalDateTime.now().toLocalTime(),
                        "Shibuya"
                    ),
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
        )
    }

}