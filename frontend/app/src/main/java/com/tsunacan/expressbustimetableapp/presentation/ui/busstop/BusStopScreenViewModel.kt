package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import android.content.Context
import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.wear.tiles.TileService
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.data.repository.UserSettingsRepository
import com.tsunacan.expressbustimetableapp.models.TimeTable
import com.tsunacan.expressbustimetableapp.tile.MainTileService
import dagger.hilt.android.lifecycle.HiltViewModel
import dagger.hilt.android.qualifiers.ApplicationContext
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.catch
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.onStart
import kotlinx.coroutines.flow.stateIn
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class BusStopScreenViewModel @Inject constructor(
    @ApplicationContext val context: Context,
    savedStateHandle: SavedStateHandle,
    private val timeTableRepository: TimeTableRepository,
    private val userSettingsRepository: UserSettingsRepository
) : ViewModel() {

    val parentRouteId: String = savedStateHandle["parentRouteId"] ?: ""
    val stopId: String = savedStateHandle["stopId"] ?: ""

    @OptIn(ExperimentalHorologistApi::class)
    val uiState: StateFlow<BusStopScreenUiState> = timeTableRepository.getTimeTable(
        parentRouteId = parentRouteId,
        busStopId = stopId
    )
        .map<TimeTable, BusStopScreenUiState> { timeTable ->
            BusStopScreenUiState.Loaded(
                timeTable = timeTable
            )
        }
        .onStart { emit(BusStopScreenUiState.Loading) }
        .catch {}
        .stateIn(viewModelScope, SharingStarted.Lazily, BusStopScreenUiState.Loading)

    fun onClickSetAsDefault(
        parentRouteId: String,
        parentRouteName: String,
        busStopId: String,
        busStopName: String
    ) {
        viewModelScope.launch {
            // Set the default bus stop to Proto DataStore
            userSettingsRepository.setDefaultBusStop(
                parentRouteId = parentRouteId,
                parentRouteName = parentRouteName,
                busStopId = busStopId,
                busStopName = busStopName
            )
            // Request an update to the tile
            TileService.getUpdater(context)
                .requestUpdate(MainTileService::class.java)
        }
    }
}
