package com.tsunacan.expressbustimetableapp.tile

import androidx.wear.protolayout.ResourceBuilders
import androidx.wear.tiles.RequestBuilders.TileRequest
import androidx.wear.tiles.TileBuilders.Tile
import androidx.wear.tiles.RequestBuilders
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.tiles.SuspendingTileService
import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import kotlinx.coroutines.flow.first

private const val RESOURCES_VERSION = "0"

@OptIn(ExperimentalHorologistApi::class)
class MainTileService : SuspendingTileService() {

    private lateinit var repo: TimeTableRepository
    private lateinit var renderer: MainTileRenderer

    override fun onCreate() {
        super.onCreate()
        repo = TimeTableRepository()
        renderer = MainTileRenderer(this)
    }

    override suspend fun resourcesRequest(
        requestParams: RequestBuilders.ResourcesRequest
    ) = resources(requestParams)

    override suspend fun tileRequest(requestParams: TileRequest): Tile {
        // TODO Initialize state in onCreate after set up Proto DataStore
        val timeTable = repo.getTimeTable("test", "test").first()
        val mainTileState = MainTileState(
            parentRouteName = "test",
            stopName = "test",
            timeTable = timeTable
        )
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
