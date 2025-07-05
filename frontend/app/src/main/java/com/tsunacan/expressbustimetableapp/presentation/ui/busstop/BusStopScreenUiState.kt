package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.models.Timetable

@ExperimentalHorologistApi
sealed class BusStopScreenUiState {

    data object Loading : BusStopScreenUiState()

    data class Loaded(
        val timetable: Timetable,
    ) : BusStopScreenUiState()

    data object Failed : BusStopScreenUiState()
}
