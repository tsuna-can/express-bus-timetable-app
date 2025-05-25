package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.data.datasource.RemoteDataSource
import com.tsunacan.expressbustimetableapp.data.mapper.BusStopMapper
import com.tsunacan.expressbustimetableapp.models.BusStop
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject

interface BusStopRepository {
    fun getBusStopList(
        parentRouteId: String
    ): Flow<List<BusStop>>
}

class BusStopRepositoryImpl @Inject constructor(
    private val remoteDataSource: RemoteDataSource,
    private val busStopMapper: BusStopMapper
) : BusStopRepository {

    override fun getBusStopList(
        parentRouteId: String
    ): Flow<List<BusStop>> {
        return flow {
            val busStopsApiModel = remoteDataSource.getBusStopList(
                parentRouteId = parentRouteId
            )
            emit(busStopMapper.mapToButStopList(busStopsApiModel))
        }
    }
}
