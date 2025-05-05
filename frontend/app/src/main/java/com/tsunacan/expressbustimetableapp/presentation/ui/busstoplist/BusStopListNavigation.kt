package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import androidx.navigation.NavController
import androidx.navigation.NavOptions
import com.tsunacan.expressbustimetableapp.presentation.ui.Screen

fun NavController.navigateToBusStopList(
    parentRouteId: String? = null,
    navOptions: NavOptions? = null,
) {
    navigate(route = Screen.BusStopList.route + "/${parentRouteId}", navOptions)
}