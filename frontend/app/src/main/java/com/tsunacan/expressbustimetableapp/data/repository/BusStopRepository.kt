package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.data.datasource.RemoteDataSource
import com.tsunacan.expressbustimetableapp.data.mapper.BusStopMapper
import com.tsunacan.expressbustimetableapp.models.BusStop
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.map
import javax.inject.Inject

interface BusStopRepository {
    fun getBusStopList(): Flow<List<BusStop>>
}

class BusStopRepositoryImpl @Inject constructor(
    private val remoteDataSource: RemoteDataSource,
    private val busStopMapper: BusStopMapper
) : BusStopRepository {

    override fun getBusStopList(): Flow<List<BusStop>> {
        return remoteDataSource.getBusStopList().map {
            busStopMapper.map(it)
        }
    }
}
