package com.tsunacan.expressbustimetableapp.tile

import androidx.lifecycle.lifecycleScope
import androidx.wear.protolayout.ResourceBuilders
import androidx.wear.tiles.RequestBuilders.TileRequest
import androidx.wear.tiles.TileBuilders.Tile
import androidx.wear.tiles.RequestBuilders
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.tiles.SuspendingTileService
import com.tsunacan.expressbustimetableapp.data.repository.UserSettingsRepository
import com.tsunacan.expressbustimetableapp.domain.GetUpcomingTimeTableUseCase
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.flow.stateIn
import java.time.Duration
import java.time.LocalTime
import javax.inject.Inject

private const val RESOURCES_VERSION = "0"

@OptIn(ExperimentalHorologistApi::class)
@AndroidEntryPoint
class MainTileService : SuspendingTileService() {

    @Inject
    lateinit var userSettingsRepository: UserSettingsRepository

    @Inject
    lateinit var getUpcomingTimeTableUseCase: GetUpcomingTimeTableUseCase

    private lateinit var renderer: MainTileRenderer

    override fun onCreate() {
        super.onCreate()
        renderer = MainTileRenderer(this)
    }

    override suspend fun resourcesRequest(
        requestParams: RequestBuilders.ResourcesRequest
    ) = resources(requestParams)

    override suspend fun tileRequest(requestParams: TileRequest): Tile {
        val defaultBusStop = userSettingsRepository.defaultBusStop.stateIn(lifecycleScope).first()

        val timeTable =
            getUpcomingTimeTableUseCase(defaultBusStop.parentRouteId, defaultBusStop.busStopId)

        val mainTileState = MainTileState(
            parentRouteName = timeTable.parentRouteName,
            stopName = timeTable.stopName,
            timeTableEntryList = timeTable.timeTableEntryList
        )

        // Set the freshness interval based on the departure time of the next bus
        val freshnessIntervalMillis = if (mainTileState.timeTableEntryList.isEmpty()) {
            Duration.ofHours(1).toMillis() // If there are no buses, refresh every hour
        } else {
            calculateTimeUntilDepartureInMs(
                mainTileState.timeTableEntryList.first().departureTime
            )
        }
        renderer.setFreshnessIntervalMillis(freshnessIntervalMillis)

        return renderer.renderTimeline(mainTileState, requestParams)
    }
}

private fun resources(
    requestParams: RequestBuilders.ResourcesRequest
): ResourceBuilders.Resources {
    return ResourceBuilders.Resources.Builder()
        .setVersion(RESOURCES_VERSION)
        .build()
}

/**
 *  Calculate the freshness interval for the tile based on the departure time of the next bus.
 *
 *  @param departureTime departure time
 *  @return time until departure in milliseconds
 */
private fun calculateTimeUntilDepartureInMs(departureTime: LocalTime): Long {
    val currentTime = LocalTime.now()
    val duration = Duration.between(currentTime, departureTime)

    return duration.toMillis()
}
