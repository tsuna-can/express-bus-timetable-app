package com.tsunacan.expressbustimetableapp.presentation.ui.busstoplist

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.compose.collectAsStateWithLifecycle
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
    viewModel: BusStopListViewModel = hiltViewModel(),
) {
    val busStopListState by viewModel.uiState.collectAsStateWithLifecycle()
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
            when (busStopListState) {
                is BusStopListScreenUiState.Loaded -> {
                    val busStopList =
                        (busStopListState as BusStopListScreenUiState.Loaded).busStopList
                    busStopList.forEach { busStopUiModel ->
                        item {
                            BusStopChip(
                                busStopUiModel = busStopUiModel,
                                navigationToBusStop = navigationToBusStop,
                                modifier = contentModifier
                            )
                        }
                    }
                }

                BusStopListScreenUiState.Loading -> {
                    item {}
                }

                BusStopListScreenUiState.Failed -> {
                    item {}
                }
            }
        }
    }
}

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun BusStopChip(
    busStopUiModel: BusStopUiModel,
    navigationToBusStop: (String, String) -> Unit,
    modifier: Modifier = Modifier,
) {
    Chip(
        modifier = modifier,
        onClick = { navigationToBusStop(busStopUiModel.parentRouteId, busStopUiModel.stopId) },
        label = {
            Text(
                text = busStopUiModel.parentRouteName,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        },
        secondaryLabel = {
            Text(
                text = busStopUiModel.stopName,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
    )
}
