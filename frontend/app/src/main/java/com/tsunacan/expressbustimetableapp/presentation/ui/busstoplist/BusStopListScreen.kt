package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.wear.compose.material.Chip
import androidx.wear.compose.material.Text
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.compose.layout.ScalingLazyColumn
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults.ItemType
import com.google.android.horologist.compose.layout.ScreenScaffold
import com.google.android.horologist.compose.layout.rememberResponsiveColumnState
import com.tsunacan.expressbustimetableapp.R
import com.tsunacan.expressbustimetableapp.models.BusStop
import com.tsunacan.expressbustimetableapp.presentation.ui.common.ErrorScreen
import com.tsunacan.expressbustimetableapp.presentation.ui.common.LoadingIndicator

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun BusStopListScreen(
    parentRouteId: String,
    navigationToBusStop: (String, String) -> Unit,
    modifier: Modifier = Modifier,
    viewModel: BusStopListScreenViewModel = hiltViewModel(),
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()

    ScreenScaffold {
        when (uiState) {
            is BusStopListScreenUiState.Loaded -> {
                BusStopListScreen(
                    busStopList = (uiState as BusStopListScreenUiState.Loaded).busStopList,
                    navigationToBusStop = navigationToBusStop,
                    modifier = modifier,
                )
            }

            BusStopListScreenUiState.Loading -> {
                LoadingIndicator()
            }

            BusStopListScreenUiState.Failed -> {
                ErrorScreen()
            }
        }
    }
}

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun BusStopListScreen(
    busStopList: List<BusStop>,
    navigationToBusStop: (String, String) -> Unit,
    modifier: Modifier = Modifier,
) {
    val contentModifier = modifier
        .fillMaxWidth()
        .padding(bottom = 8.dp)

    val listState = rememberResponsiveColumnState(
        contentPadding = ScalingLazyColumnDefaults.padding(
            first = ItemType.Chip,
            last = ItemType.Chip,
        ),
    )

    val parentRouteName = busStopList.firstOrNull()?.parentRouteName ?: "Unknown Route"

    ScalingLazyColumn(
        columnState = listState,
    ) {
        item {
            Text(
                text = parentRouteName,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis,
            )
        }
        if (busStopList.isEmpty()) {
            item {
                Text(
                    text = stringResource(R.string.no_bus_stop),
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis,
                )
            }
        } else {
            busStopList.forEach { busStop ->
                item {
                    BusStopChip(
                        busStop = busStop,
                        navigationToBusStop = navigationToBusStop,
                        modifier = contentModifier
                    )
                }
            }
        }
    }
}

@Composable
fun BusStopChip(
    busStop: BusStop,
    navigationToBusStop: (String, String) -> Unit,
    modifier: Modifier = Modifier,
) {
    Chip(
        modifier = modifier,
        onClick = { navigationToBusStop(busStop.parentRouteId, busStop.stopId) },
        label = {
            Text(
                text = busStop.stopName,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
    )
}
