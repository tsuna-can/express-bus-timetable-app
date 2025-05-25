package com.tsunacan.expressbustimetableapp.presentation.ui.parentroutelist

import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.models.ParentRoute

@ExperimentalHorologistApi
sealed class ParentRouteListScreenUiState {

    data object Loading : ParentRouteListScreenUiState()

    data class Loaded(
        val parentRouteList: List<ParentRoute>,
    ) : ParentRouteListScreenUiState()

    data object Failed : ParentRouteListScreenUiState()
}
