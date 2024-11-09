package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import com.google.android.horologist.annotations.ExperimentalHorologistApi

@ExperimentalHorologistApi
data class BusStopUiModel(
    val parentRouteId: String,
    val stopId: String,
    val parentRouteName: String,
    val stopName: String,
)
