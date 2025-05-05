package com.tsunacan.expressbustimetableapp.presentation.ui

sealed class Screen(
    val route: String,
) {
    data object ParentRouteList : Screen("parentRouteList")
    data object BusStop : Screen("busStop")
    data object BusStopList : Screen("busStopList")
}
