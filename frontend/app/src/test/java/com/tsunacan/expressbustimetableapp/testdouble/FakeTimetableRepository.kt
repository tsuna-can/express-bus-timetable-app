package com.tsunacan.expressbustimetableapp.testdouble

import com.tsunacan.expressbustimetableapp.data.repository.TimetableRepository
import com.tsunacan.expressbustimetableapp.models.Timetable
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow

class FakeTimetableRepository : TimetableRepository {

    private val timetables = mutableMapOf<Pair<String, String>, Timetable>()

    fun addTimetable(parentRouteId: String, busStopId: String, timetable: Timetable) {
        timetables[parentRouteId to busStopId] = timetable
    }

    override fun getTimetable(parentRouteId: String, busStopId: String): Flow<Timetable> {
        val timetable = timetables[parentRouteId to busStopId]
            ?: throw IllegalArgumentException("No timetable found for route $parentRouteId and stop $busStopId")
        return flow { emit(timetable) }
    }
}
