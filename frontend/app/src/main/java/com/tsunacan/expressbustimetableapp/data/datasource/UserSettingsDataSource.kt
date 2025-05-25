package com.tsunacan.expressbustimetableapp.data.datasource

import android.util.Log
import androidx.datastore.core.DataStore
import com.tsunacan.expressbustimetableapp.DefaultBusStop
import com.tsunacan.expressbustimetableapp.copy
import java.io.IOException
import javax.inject.Inject

class UserSettingsDataSource @Inject constructor(
    private val userSettings: DataStore<DefaultBusStop>,
) {

    val defaultBusStop = userSettings.data

    suspend fun setDefaultBusStop(
        parentRouteId: String,
        parentRouteName: String,
        busStopId: String,
        busStopName: String
    ) {
        try {
            userSettings.updateData {
                it.copy {
                    this.parentRouteId = parentRouteId
                    this.parentRouteName = parentRouteName
                    this.busStopId = busStopId
                    this.busStopName = busStopName
                }
            }
        } catch (ioException: IOException) {
            Log.e("Proto DataStore", "Failed to set default bus stop", ioException)
        }
    }
}
