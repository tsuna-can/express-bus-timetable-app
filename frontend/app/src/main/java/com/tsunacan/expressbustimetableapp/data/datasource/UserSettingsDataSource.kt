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

    suspend fun setDefaultBusStop(busStop: DefaultBusStop) {
        try{
            userSettings.updateData {
                it.copy {
                    parentRouteId = busStop.parentRouteId
                    busStopId = busStop.busStopId
                }
            }
        } catch (ioException: IOException) {
            Log.e("Proto DataStore", "Failed to set default bus stop", ioException)
        }
    }

}
