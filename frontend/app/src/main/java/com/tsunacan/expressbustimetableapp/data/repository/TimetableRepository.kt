package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.data.datasource.RemoteDataSource
import com.tsunacan.expressbustimetableapp.data.mapper.TimetableMapper
import com.tsunacan.expressbustimetableapp.models.Timetable
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject

interface TimetableRepository {
    fun getTimetable(
        parentRouteId: String,
        busStopId: String
    ): Flow<Timetable>
}

class TimetableRepositoryImpl @Inject constructor(
    private val remoteDataSource: RemoteDataSource,
    private val timetableMapper: TimetableMapper
) : TimetableRepository {

    override fun getTimetable(
        parentRouteId: String,
        busStopId: String
    ): Flow<Timetable> {
        return flow {
            val timetable = remoteDataSource.getTimetable(parentRouteId, busStopId)
            emit(timetableMapper.mapToTimetable(timetable))
        }
    }
}
