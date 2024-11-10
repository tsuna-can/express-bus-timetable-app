package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.models.TimeTable

@ExperimentalHorologistApi
sealed class BusStopScreenUiState {

    data object Loading : BusStopScreenUiState()

    data class Loaded(
        val timeTable: TimeTable,
    ) : BusStopScreenUiState()

    data object Failed : BusStopScreenUiState()
}
