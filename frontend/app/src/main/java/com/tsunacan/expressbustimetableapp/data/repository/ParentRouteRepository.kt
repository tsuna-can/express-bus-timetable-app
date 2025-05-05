package com.tsunacan.expressbustimetableapp.data.repository

import android.util.Log
import com.tsunacan.expressbustimetableapp.data.datasource.RemoteDataSource
import com.tsunacan.expressbustimetableapp.data.mapper.ParentRouteMapper
import com.tsunacan.expressbustimetableapp.models.ParentRoute
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject

interface ParentRouteRepository {
    fun getParentRouteList(): Flow<List<ParentRoute>>
}

class ParentRouteRepositoryImpl @Inject constructor(
    private val remoteDataSource: RemoteDataSource,
    private val parentRouteMapper: ParentRouteMapper
) : ParentRouteRepository {

    override fun getParentRouteList(): Flow<List<ParentRoute>> {
        return flow {
            try {
                val parentRoutesApiModel= remoteDataSource.getParentRouteList()
                emit(parentRouteMapper.map(parentRoutesApiModel.parentRoutes))
            } catch (e: Exception) {
                Log.e("ParentRouteRepository", "Error fetching parent route list", e)
                emit(emptyList())
            }
        }
    }
}