package com.tsunacan.expressbustimetableapp.data.datasource

import androidx.compose.ui.util.trace
import com.tsunacan.expressbustimetableapp.BuildConfig
import com.tsunacan.expressbustimetableapp.data.model.BusStopsApiModel
import com.tsunacan.expressbustimetableapp.data.model.ParentRoutesApiModel
import com.tsunacan.expressbustimetableapp.data.model.TimeTableApiModel
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.serialization.json.Json
import okhttp3.Call
import okhttp3.MediaType.Companion.toMediaType
import retrofit2.Retrofit
import retrofit2.converter.kotlinx.serialization.asConverterFactory
import retrofit2.http.GET
import retrofit2.http.Query
import javax.inject.Inject

interface NetworkDataSource {
    suspend fun getParentRouteList(): ParentRoutesApiModel

    suspend fun getBusStopList(
        parentRouteId: String
    ): BusStopsApiModel

    suspend fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): TimeTableApiModel

}

private interface NetworkApi {
    @GET(value = "parent-routes")
    suspend fun getParentRouteList(): ParentRoutesApiModel

    @GET(value = "bus-stops")
    suspend fun getBusStopList(
        @Query("parent-route-id") parentRouteId: String
    ): BusStopsApiModel

    @GET(value = "timetable")
    suspend fun getTimeTable(
        @Query("parent-route-id") parentRouteId: String,
        @Query("bus-stop-id") busStopId: String
    ): TimeTableApiModel
}

class RemoteDataSource @Inject constructor(
    private val okhttpCallFactory: dagger.Lazy<Call.Factory>,
    private val networkJson: Json,
    private val ioDispatcher: CoroutineDispatcher
): NetworkDataSource{

    private val networkApi = trace("RetrofitNetwork") {
        Retrofit.Builder()
            .baseUrl(BuildConfig.BASE_URL)
            // We use callFactory lambda here with dagger.Lazy<Call.Factory>
            // to prevent initializing OkHttp on the main thread.
            .callFactory { okhttpCallFactory.get().newCall(it) }
            .addConverterFactory(
                networkJson.asConverterFactory("application/json".toMediaType()),
            )
            .build()
            .create(NetworkApi::class.java)
    }

    override suspend fun getParentRouteList(): ParentRoutesApiModel {
        return networkApi.getParentRouteList()
    }

    override suspend fun getBusStopList(
        parentRouteId: String
    ): BusStopsApiModel {
        return networkApi.getBusStopList(parentRouteId)
    }

    override suspend fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): TimeTableApiModel {
        return networkApi.getTimeTable(parentRouteId, busStopId)
    }
}
