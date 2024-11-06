package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import androidx.navigation.NavController
import androidx.navigation.NavOptions
import com.tsunacan.expressbustimetableapp.presentation.ui.Screen

fun NavController.navigateToBusStop(
    parentRouteId: String? = null,
    stopId: String? = null,
    navOptions: NavOptions? = null,
) {
    navigate(route = Screen.BusStop.route + "/${parentRouteId}/${stopId}", navOptions)
}
