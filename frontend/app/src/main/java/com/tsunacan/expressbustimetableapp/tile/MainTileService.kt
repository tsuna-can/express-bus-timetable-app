package com.tsunacan.expressbustimetableapp.tile

import androidx.wear.protolayout.ResourceBuilders
import androidx.wear.tiles.RequestBuilders.TileRequest
import androidx.wear.tiles.TileBuilders.Tile
import androidx.wear.tiles.RequestBuilders
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.tiles.SuspendingTileService
import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter

private const val RESOURCES_VERSION = "0"

@OptIn(ExperimentalHorologistApi::class)
class MainTileService : SuspendingTileService() {

    private lateinit var renderer: MainTileRenderer

    override fun onCreate() {
        super.onCreate()
        renderer = MainTileRenderer(this)
    }

    override suspend fun resourcesRequest(
        requestParams: RequestBuilders.ResourcesRequest
    ) = resources(requestParams)

    override suspend fun tileRequest(requestParams: TileRequest): Tile {
        val mainTileState = generateDummyTileState()
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

private fun generateDummyTileState(): MainTileState {
    return MainTileState(
        parentRouteName = "Tokyo",
        stopName = "Shinjuku",
        timeTable = listOf(
            DepartureTimeAndDestination(LocalDateTime.now().plusMinutes(10).toLocalTime(), "Shibuya"),
            DepartureTimeAndDestination(LocalDateTime.now().plusMinutes(10).toLocalTime(), "Shibuya"),
            DepartureTimeAndDestination(LocalDateTime.now().plusMinutes(10).toLocalTime(), "Shibuya"),
        )
    )
}
