package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import android.content.Context
import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.wear.tiles.TileService
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.data.repository.UserSettingsRepository
import com.tsunacan.expressbustimetableapp.domain.GetDaySpecificTimetableUseCase
import com.tsunacan.expressbustimetableapp.tile.MainTileService
import dagger.hilt.android.lifecycle.HiltViewModel
import dagger.hilt.android.qualifiers.ApplicationContext
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.flow.stateIn
import kotlinx.coroutines.launch
import java.time.LocalDate
import javax.inject.Inject

@HiltViewModel
class BusStopScreenViewModel @Inject constructor(
    @ApplicationContext val context: Context,
    savedStateHandle: SavedStateHandle,
    private val getDaySpecificTimetableUseCase: GetDaySpecificTimetableUseCase,
    private val userSettingsRepository: UserSettingsRepository
) : ViewModel() {

    val parentRouteId: String = savedStateHandle["parentRouteId"] ?: ""
    val stopId: String = savedStateHandle["stopId"] ?: ""

    @OptIn(ExperimentalHorologistApi::class)
    val uiState: StateFlow<BusStopScreenUiState> = flow {
        emit(BusStopScreenUiState.Loading)
        try {
            val timetable = getDaySpecificTimetableUseCase.invoke(
                parentRouteId = parentRouteId,
                busStopId = stopId,
                today = LocalDate.now()
            )
            emit(BusStopScreenUiState.Loaded(timetable = timetable))
        } catch (e: Exception) {
            emit(BusStopScreenUiState.Failed)
        }
    }
        .stateIn(viewModelScope, SharingStarted.Lazily, BusStopScreenUiState.Loading)

    fun onClickSetForTile(
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
