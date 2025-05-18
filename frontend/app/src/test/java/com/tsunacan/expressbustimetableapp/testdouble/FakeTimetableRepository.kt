package com.tsunacan.expressbustimetableapp.testdouble

import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.models.TimeTable
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow

class FakeTimetableRepository : TimeTableRepository {

    private val timeTables = mutableMapOf<Pair<String, String>, TimeTable>()

    fun addTimeTable(parentRouteId: String, busStopId: String, timeTable: TimeTable) {
        timeTables[parentRouteId to busStopId] = timeTable
    }

    override fun getTimeTable(parentRouteId: String, busStopId: String): Flow<TimeTable> {
        val timeTable = timeTables[parentRouteId to busStopId]
            ?: throw IllegalArgumentException("No timetable found for route $parentRouteId and stop $busStopId")
        return flow { emit(timeTable) }
    }
}
