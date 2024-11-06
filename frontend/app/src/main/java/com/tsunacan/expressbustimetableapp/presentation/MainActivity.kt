/* While this template provides a good starting point for using Wear Compose, you can always
 * take a look at https://github.com/android/wear-os-samples/tree/main/ComposeStarter to find the
 * most up to date changes to the libraries and their usages.
 */

package com.tsunacan.expressbustimetableapp.presentation

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.core.splashscreen.SplashScreen.Companion.installSplashScreen
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
import com.google.android.horologist.compose.layout.AppScaffold
import com.tsunacan.expressbustimetableapp.presentation.theme.ExpressBusTimeTableAppTheme
import com.tsunacan.expressbustimetableapp.presentation.ui.Screen
import com.tsunacan.expressbustimetableapp.presentation.ui.busstop.BusStopScreen
import com.tsunacan.expressbustimetableapp.presentation.ui.busstop.navigateToBusStop
import com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist.BusStopListScreen
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        installSplashScreen()

        super.onCreate(savedInstanceState)

        setTheme(android.R.style.Theme_DeviceDefault)

        setContent {
            WearApp()
        }
    }
}

@Composable
fun WearApp() {
    val swipeToDismissBoxState = rememberSwipeToDismissBoxState()
    val navHostState =
        rememberSwipeDismissableNavHostState(swipeToDismissBoxState = swipeToDismissBoxState)
    val navController = rememberSwipeDismissableNavController()

    ExpressBusTimeTableAppTheme {
        AppScaffold {
            SwipeDismissableNavHost(
                startDestination = Screen.BusStopList.route,
                navController = navController,
                state = navHostState,
            ) {
                composable(
                    route = Screen.BusStopList.route,
                ) {
                    BusStopListScreen(
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
    WearApp()
}
