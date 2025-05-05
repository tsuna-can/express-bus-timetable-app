/* While this template provides a good starting point for using Wear Compose, you can always
 * take a look at https://github.com/android/wear-os-samples/tree/main/ComposeStarter to find the
 * most up to date changes to the libraries and their usages.
 */

package com.tsunacan.expressbustimetableapp.presentation

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.core.splashscreen.SplashScreen.Companion.installSplashScreen
import com.tsunacan.expressbustimetableapp.presentation.ui.Screen
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        installSplashScreen()

        super.onCreate(savedInstanceState)

        setTheme(android.R.style.Theme_DeviceDefault)

        // Get destination, parentRouteId, busStopId from the intent
        val destination = intent.getStringExtra("destination") ?: ""
        val parentRouteId = intent.getStringExtra("parentRouteId") ?: ""
        val busStopId = intent.getStringExtra("busStopId") ?: ""

        // If the destination is a busStop, need to pass the parentRouteId and busStopId
        val initialRoute = when {
            destination == Screen.BusStop.route -> Screen.BusStop.route + "/$parentRouteId/$busStopId"
            else -> Screen.ParentRouteList.route
        }

        setContent {
            WearApp(initialRoute)
        }
    }
}
