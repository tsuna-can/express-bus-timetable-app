package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

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
import androidx.wear.compose.material.Button
import androidx.wear.compose.material.Text
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.compose.layout.ScalingLazyColumn
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults.ItemType
import com.google.android.horologist.compose.layout.ScreenScaffold
import com.google.android.horologist.compose.layout.rememberResponsiveColumnState
import com.google.android.horologist.compose.material.Chip
import com.tsunacan.expressbustimetableapp.R
import com.tsunacan.expressbustimetableapp.models.TimeTable
import java.time.LocalTime

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun BusStopScreen(
    parentRouteId: String,
    stopId: String,
    modifier: Modifier = Modifier,
    viewModel: BusStopScreenViewModel = hiltViewModel(),
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()

    ScreenScaffold() {
        when (uiState) {
            is BusStopScreenUiState.Loaded -> {
                BusStopScreen(
                    timeTable = (uiState as BusStopScreenUiState.Loaded).timeTable,
                    onClickSetAsDefault = viewModel::onClickSetAsDefault,
                    modifier = modifier,
                )
            }

            BusStopScreenUiState.Loading -> {
            }

            BusStopScreenUiState.Failed -> {
            }
        }
    }
}

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun BusStopScreen(
    timeTable: TimeTable,
    onClickSetAsDefault: () -> Unit,
    modifier: Modifier = Modifier,
) {
    val contentModifier = Modifier
        .fillMaxWidth()
        .padding(bottom = 8.dp)

    val listState = rememberResponsiveColumnState(
        contentPadding = ScalingLazyColumnDefaults.padding(
            first = ItemType.Text,
            last = ItemType.Chip,
        ),
    )

    val timeTableEntryList = timeTable.timeTableEntryList

    ScalingLazyColumn(
        columnState = listState,
    ) {
        item {
            Text(
                text = timeTable.parentRouteName,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
        item {
            Text(
                text = timeTable.stopName,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
        item {
            Button(
                onClick = onClickSetAsDefault,
                modifier = contentModifier,
            ) {
                Text(
                    text = stringResource(R.string.set_as_default),
                )
            }
        }
        timeTableEntryList.forEach { timeTableEntry ->
            item {
                TimeTableEntryChip(
                    modifier = contentModifier,
                    departureTime = timeTableEntry.departureTime,
                    destination = timeTableEntry.destination,
                )
            }
        }
    }
}

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun TimeTableEntryChip(
    departureTime: LocalTime,
    destination: String,
    modifier: Modifier = Modifier,
) {
    Chip(
        modifier = modifier,
        onClick = {},
        label = {
            Text(
                text = departureTime.toString(),
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        },
        secondaryLabel = {
            Text(
                text = destination,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
    )
}
