package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import com.google.android.horologist.annotations.ExperimentalHorologistApi

@ExperimentalHorologistApi
sealed class BusStopListScreenUiState {

    data object Loading : BusStopListScreenUiState()

    data class Loaded(
        val busStopList: List<BusStopUiModel>,
    ) : BusStopListScreenUiState()

    data object Failed : BusStopListScreenUiState()
}
