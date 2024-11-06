package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.wear.compose.material.Chip
import androidx.wear.compose.material.Text
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.compose.layout.ScalingLazyColumn
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults
import com.google.android.horologist.compose.layout.ScreenScaffold
import com.google.android.horologist.compose.layout.rememberResponsiveColumnState
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults.ItemType

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun BusStopListScreen(
    navigationToBusStop: (String, String) -> Unit,
    modifier: Modifier = Modifier,
) {
    val listState = rememberResponsiveColumnState(
        contentPadding = ScalingLazyColumnDefaults.padding(
            first = ItemType.Chip,
            last = ItemType.Chip,
        ),
    )

    ScreenScaffold(scrollState = listState) {

        val contentModifier = Modifier
            .fillMaxWidth()
            .padding(bottom = 8.dp)

        ScalingLazyColumn(
            columnState = listState,
        ) {
            item { BusStopChip("routeId1", "stopId1", navigationToBusStop, contentModifier) }
            item { BusStopChip("routeId2", "stopId2", navigationToBusStop, contentModifier) }
            item { BusStopChip("routeId3", "stopId3", navigationToBusStop, contentModifier) }
        }
    }
}

@Composable
fun BusStopChip(
    // TODO pass parameters as data class
    parentRouteId: String,
    stopId: String,
    navigationToBusStop: (String, String) -> Unit,
    modifier: Modifier = Modifier,
) {
    Chip(
        modifier = modifier,
        onClick = { navigationToBusStop(parentRouteId, stopId) },
        label = {
            Text(
                text = "Tokyo express, Tokyo station",
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        },
        secondaryLabel = {
            Text(
                text = "inbound",
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
    )
}
