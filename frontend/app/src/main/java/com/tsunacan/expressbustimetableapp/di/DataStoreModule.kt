package com.tsunacan.expressbustimetableapp.di

import android.content.Context
import androidx.datastore.dataStoreFile
import androidx.datastore.core.DataStore
import androidx.datastore.core.DataStoreFactory
import com.tsunacan.expressbustimetableapp.DefaultBusStop
import com.tsunacan.expressbustimetableapp.data.UserSettingsSerializer
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.components.SingletonComponent
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object DataStoreModule {

    @Provides
    @Singleton
    internal fun providesUserSettingsDataStore(
        @ApplicationContext context: Context,
        userSettingsSerializer: UserSettingsSerializer,
    ): DataStore<DefaultBusStop> =
        DataStoreFactory.create(
            serializer = userSettingsSerializer,
        ){
            context.dataStoreFile("user_settings.pb")
        }
}