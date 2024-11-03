package com.tsunacan.expressbustimetableapp.data.datasource

import androidx.compose.ui.util.trace
import com.tsunacan.expressbustimetableapp.BuildConfig
import com.tsunacan.expressbustimetableapp.data.model.TimeTableApiModel
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.flow.flowOn
import kotlinx.serialization.json.Json
import okhttp3.Call
import okhttp3.MediaType.Companion.toMediaType
import retrofit2.Retrofit
import retrofit2.converter.kotlinx.serialization.asConverterFactory
import retrofit2.http.GET
import retrofit2.http.Query
import javax.inject.Inject

private interface NetworkApi {
    @GET(value = "timetable")
    suspend fun getTimeTable(
        @Query("parentRouteId") parentRouteId: String,
        @Query("busStopId") busStopId: String
    ): TimeTableApiModel
}

class RemoteDataSource @Inject constructor(
    private val okhttpCallFactory: dagger.Lazy<Call.Factory>,
    private val networkJson: Json,
    private val ioDispatcher: CoroutineDispatcher
) {

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

    fun getTimeTable(
        parentRouteId: String,
        busStopId: String
    ): Flow<TimeTableApiModel> {
        return flow { emit(networkApi.getTimeTable(parentRouteId, busStopId)) }.flowOn(ioDispatcher)
    }
}
