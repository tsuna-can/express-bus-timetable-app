package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.flowOf
import kotlinx.coroutines.flow.stateIn
import javax.inject.Inject

@ExperimentalHorologistApi
@HiltViewModel
class BusStopListViewModel @Inject constructor() : ViewModel() {

    val uiState: StateFlow<BusStopListScreenUiState> =
        flowOf(
            BusStopListScreenUiState.Loaded(
                listOf(
                    BusStopUiModel(
                        parentRouteId = "Route Id 1",
                        parentRouteName = "Route 1",
                        stopId = "Stop Id 1",
                        stopName = "Bus Stop 1"
                    ),
                    BusStopUiModel(
                        parentRouteId = "Route Id 2",
                        parentRouteName = "Route 2",
                        stopId = "Stop Id 2",
                        stopName = "Bus Stop 2"
                    ),
                    BusStopUiModel(
                        parentRouteId = "Route Id 3",
                        parentRouteName = "Route 3",
                        stopId = "Stop Id 3",
                        stopName = "Bus Stop 3"
                    ),
                )
            )
        )
            .stateIn(
                scope = viewModelScope,
                started = SharingStarted.WhileSubscribed(5_000),
                initialValue = BusStopListScreenUiState.Loading
            )

}