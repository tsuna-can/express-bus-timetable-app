package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.data.datasource.RemoteDataSource
import com.tsunacan.expressbustimetableapp.data.mapper.TimeTableMapper
import com.tsunacan.expressbustimetableapp.models.TimeTable
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.map
import javax.inject.Inject

interface TimeTableRepository {
    suspend fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): Flow<TimeTable>
}

class TimeTableRepositoryImpl @Inject constructor(
    private val remoteDataSource: RemoteDataSource,
    private val timeTableMapper: TimeTableMapper
) : TimeTableRepository {

    override suspend fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): Flow<TimeTable> {
        return remoteDataSource.getTimeTable(parentRouteId, busStopId).map{
            timeTableMapper.map(it)
        }
    }
}
