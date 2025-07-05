package com.tsunacan.expressbustimetableapp.di

import androidx.datastore.core.DataStore
import com.tsunacan.expressbustimetableapp.DefaultBusStop
import com.tsunacan.expressbustimetableapp.data.datasource.RemoteDataSource
import com.tsunacan.expressbustimetableapp.data.datasource.UserSettingsDataSource
import com.tsunacan.expressbustimetableapp.data.mapper.BusStopMapper
import com.tsunacan.expressbustimetableapp.data.mapper.ParentRouteMapper
import com.tsunacan.expressbustimetableapp.data.mapper.TimetableMapper
import com.tsunacan.expressbustimetableapp.data.repository.BusStopRepository
import com.tsunacan.expressbustimetableapp.data.repository.BusStopRepositoryImpl
import com.tsunacan.expressbustimetableapp.data.repository.ParentRouteRepository
import com.tsunacan.expressbustimetableapp.data.repository.ParentRouteRepositoryImpl
import com.tsunacan.expressbustimetableapp.data.repository.TimetableRepository
import com.tsunacan.expressbustimetableapp.data.repository.TimetableRepositoryImpl
import com.tsunacan.expressbustimetableapp.data.repository.UserSettingsRepository
import com.tsunacan.expressbustimetableapp.data.repository.UserSettingsRepositoryImpl
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.coroutines.Dispatchers
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
class DataModule {

    @Provides
    @Singleton
    fun parentRouteRepository(
        parentRouteRepositoryImpl: ParentRouteRepositoryImpl
    ): ParentRouteRepository = parentRouteRepositoryImpl

    @Provides
    @Singleton
    fun parentRouteRepositoryImpl(
        remoteDataSource: RemoteDataSource,
        parentRouteMapper: ParentRouteMapper
    ) = ParentRouteRepositoryImpl(remoteDataSource, parentRouteMapper)

    @Provides
    @Singleton
    fun busStopRepository(
        busStopRepositoryImpl: BusStopRepositoryImpl
    ): BusStopRepository = busStopRepositoryImpl

    @Provides
    @Override
    fun busStopRepositoryImpl(
        remoteDataSource: RemoteDataSource,
        busStopMapper: BusStopMapper
    ) = BusStopRepositoryImpl(remoteDataSource, busStopMapper)

    @Provides
    @Singleton
    fun timetableRepository(
        timetableRepositoryImpl: TimetableRepositoryImpl
    ): TimetableRepository = timetableRepositoryImpl

    @Provides
    @Singleton
    fun timetableRepositoryImpl(
        remoteDataSource: RemoteDataSource,
        timetableMapper: TimetableMapper
    ) = TimetableRepositoryImpl(remoteDataSource, timetableMapper)

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

    @Provides
    fun parentRouteMapper() = ParentRouteMapper

    @Provides
    fun busStopMapper() = BusStopMapper

    @Provides
    fun timetableMapper() = TimetableMapper

    @Provides
    fun ioDispatcher(): CoroutineDispatcher = Dispatchers.IO
}
