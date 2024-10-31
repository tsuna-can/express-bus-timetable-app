package com.tsunacan.expressbustimetableapp.data.repository

import com.tsunacan.expressbustimetableapp.DefaultBusStop
import com.tsunacan.expressbustimetableapp.data.datasource.UserSettingsDataSource
import javax.inject.Inject

interface UserSettingsRepository {
    suspend fun setDefaultBusStop(defaultBusStop: DefaultBusStop)
}

class UserSettingsRepositoryImpl @Inject constructor(
    private val userSettingsDataSource: UserSettingsDataSource
) : UserSettingsRepository {

    override suspend fun setDefaultBusStop(defaultBusStop: DefaultBusStop) {
        userSettingsDataSource.setDefaultBusStop(defaultBusStop)
    }
}
