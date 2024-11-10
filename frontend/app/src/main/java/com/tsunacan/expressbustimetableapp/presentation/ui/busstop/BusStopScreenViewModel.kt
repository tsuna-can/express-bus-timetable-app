package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.data.repository.TimeTableRepository
import com.tsunacan.expressbustimetableapp.models.TimeTable
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.catch
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.onStart
import kotlinx.coroutines.flow.stateIn
import javax.inject.Inject

@HiltViewModel
class BusStopScreenViewModel @Inject constructor(
    savedStateHandle: SavedStateHandle,
    private val timeTableRepository: TimeTableRepository
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
}
