package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.data.repository.BusStopRepository
import com.tsunacan.expressbustimetableapp.models.BusStop
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.catch
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.onStart
import kotlinx.coroutines.flow.stateIn
import javax.inject.Inject

@ExperimentalHorologistApi
@HiltViewModel
class BusStopListScreenViewModel @Inject constructor(
    savedStateHandle: SavedStateHandle,
    private val busStopRepository: BusStopRepository
) : ViewModel() {

    val parentRouteId :String = savedStateHandle["parentRouteId"] ?: ""

    val uiState: StateFlow<BusStopListScreenUiState> = busStopRepository.getBusStopList(
        parentRouteId = parentRouteId
    )
        .map<List<BusStop>, BusStopListScreenUiState> { busStops ->
            BusStopListScreenUiState.Loaded(busStops)
        }
        .onStart { emit(BusStopListScreenUiState.Loading) }
        .catch {}
        .stateIn(viewModelScope, SharingStarted.Lazily, BusStopListScreenUiState.Loading)

}
