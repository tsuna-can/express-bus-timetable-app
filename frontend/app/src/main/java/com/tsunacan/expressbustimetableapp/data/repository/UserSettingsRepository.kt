package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.DefaultBusStop
import com.tsunacan.expressbustimetableapp.data.datasource.UserSettingsDataSource
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject

interface UserSettingsRepository {
    val defaultBusStop: Flow<DefaultBusStop>
    suspend fun setDefaultBusStop(
        parentRouteId: String,
        parentRouteName: String,
        busStopId: String,
        busStopName: String
    )
}

class UserSettingsRepositoryImpl @Inject constructor(
    private val userSettingsDataSource: UserSettingsDataSource
) : UserSettingsRepository {

    override val defaultBusStop: Flow<DefaultBusStop> = userSettingsDataSource.defaultBusStop

    override suspend fun setDefaultBusStop(
        parentRouteId: String,
        parentRouteName: String,
        busStopId: String,
        busStopName: String
    ) {
        userSettingsDataSource.setDefaultBusStop(
            parentRouteId,
            parentRouteName,
            busStopId,
            busStopName
        )
    }
}
