package com.tsunacan.expressbustimetableapp.di

import androidx.datastore.core.DataStore
import com.tsunacan.expressbustimetableapp.DefaultBusStop
import com.tsunacan.expressbustimetableapp.data.datasource.UserSettingsDataSource
import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepositoryImpl
import com.tsunacan.expressbustimetableapp.data.repository.UserSettingsRepository
import com.tsunacan.expressbustimetableapp.data.repository.UserSettingsRepositoryImpl
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
class DataModule {

    @Provides
    @Singleton
    fun timeTableRepository(): TimeTableRepository = TimeTableRepositoryImpl()

    @Provides
    @Singleton
    fun userSettingsRepository(
        userSettingsRepositoryImpl: UserSettingsRepositoryImpl
    ): UserSettingsRepository = userSettingsRepositoryImpl

    @Provides
    @Singleton
    fun userSettingsRepositoryImpl(
        userSettingsDataSource: UserSettingsDataSource,
    ) = UserSettingsRepositoryImpl(userSettingsDataSource)

    @Provides
    @Singleton
    fun userSettingsDataSource(
        dataStore: DataStore<DefaultBusStop>
    ): UserSettingsDataSource = UserSettingsDataSource(dataStore)

}