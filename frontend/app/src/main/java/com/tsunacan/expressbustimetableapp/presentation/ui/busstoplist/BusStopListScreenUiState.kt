package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.models.BusStop

@ExperimentalHorologistApi
sealed class BusStopListScreenUiState {

    data object Loading : BusStopListScreenUiState()

    data class Loaded(
        val busStopList: List<BusStop>,
    ) : BusStopListScreenUiState()

    data object Failed : BusStopListScreenUiState()
}
