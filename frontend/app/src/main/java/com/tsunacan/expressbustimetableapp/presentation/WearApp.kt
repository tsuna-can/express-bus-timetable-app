package com.tsunacan.expressbustimetableapp.presentation

import androidx.compose.runtime.Composable
import androidx.compose.ui.tooling.preview.Preview
import androidx.navigation.NavType
import androidx.navigation.navArgument
import androidx.wear.compose.foundation.rememberSwipeToDismissBoxState
import androidx.wear.compose.navigation.SwipeDismissableNavHost
import androidx.wear.compose.navigation.composable
import androidx.wear.compose.navigation.rememberSwipeDismissableNavController
import androidx.wear.compose.navigation.rememberSwipeDismissableNavHostState
import androidx.wear.tooling.preview.devices.WearDevices
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.compose.layout.AppScaffold
import com.tsunacan.expressbustimetableapp.presentation.theme.ExpressBusTimeTableAppTheme
import com.tsunacan.expressbustimetableapp.presentation.ui.Screen
import com.tsunacan.expressbustimetableapp.presentation.ui.busstop.BusStopScreen
import com.tsunacan.expressbustimetableapp.presentation.ui.busstop.navigateToBusStop
import com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist.BusStopListScreen
import com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist.navigateToBusStopList
import com.tsunacan.expressbustimetableapp.presentation.ui.parentroutelist.ParentRouteListScreen

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun WearApp(
    initialRoute: String
) {
    val swipeToDismissBoxState = rememberSwipeToDismissBoxState()
    val navHostState =
        rememberSwipeDismissableNavHostState(swipeToDismissBoxState = swipeToDismissBoxState)
    val navController = rememberSwipeDismissableNavController()

    ExpressBusTimeTableAppTheme {
        AppScaffold {
            SwipeDismissableNavHost(
                startDestination = initialRoute,
                navController = navController,
                state = navHostState,
            ) {
                composable(
                    route = Screen.ParentRouteList.route,
                ) {
                    ParentRouteListScreen(
                        navigationToBusStopList = navController::navigateToBusStopList,
                    )
                }
                composable(
                    route = Screen.BusStopList.route + "/{parentRouteId}",
                    arguments = listOf(
                        navArgument("parentRouteId") { type = NavType.StringType },
                    )
                ) { navBackStackEntry ->
                    val parentRouteId = navBackStackEntry.arguments?.getString("parentRouteId")
                    BusStopListScreen(
                        parentRouteId = parentRouteId ?: "",
                        navigationToBusStop = navController::navigateToBusStop
                    )
                }
                composable(
                    route = Screen.BusStop.route + "/{parentRouteId}/{stopId}",
                    arguments = listOf(
                        navArgument("parentRouteId") { type = NavType.StringType },
                        navArgument("stopId") { type = NavType.StringType }
                    )
                ) { navBackStackEntry ->
                    val parentRouteId = navBackStackEntry.arguments?.getString("parentRouteId")
                    val stopId = navBackStackEntry.arguments?.getString("stopId")
                    BusStopScreen(
                        parentRouteId = parentRouteId ?: "",
                        stopId = stopId ?: ""
                    )
                }
            }
        }
    }
}

@Preview(device = WearDevices.SMALL_ROUND, showSystemUi = true)
@Composable
fun DefaultPreview() {
    WearApp(
        initialRoute = Screen.BusStopList.route
    )
}
