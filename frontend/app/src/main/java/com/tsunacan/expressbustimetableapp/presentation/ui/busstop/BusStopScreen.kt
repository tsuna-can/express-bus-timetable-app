package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import android.widget.Toast
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.wear.compose.material.Button
import androidx.wear.compose.material.ButtonDefaults
import androidx.wear.compose.material.MaterialTheme
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
import com.tsunacan.expressbustimetableapp.presentation.ui.common.LoadingIndicator
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
                    onClickSetForTile = viewModel::onClickSetForTile,
                    modifier = modifier,
                )
            }

            BusStopScreenUiState.Loading -> {
                LoadingIndicator()
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
    onClickSetForTile: (String, String, String, String) -> Unit,
    modifier: Modifier = Modifier,
) {
    val contentModifier = modifier
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
            SetForTileButton(
                onClick = onClickSetForTile(
                    timeTable.parentRouteId,
                    timeTable.parentRouteName,
                    timeTable.stopId,
                    timeTable.stopName
                ),
                modifier = contentModifier,
            )
        }
        if (timeTableEntryList.isEmpty()) {
            item {
                Text(
                    text = stringResource(R.string.no_timetable_entry),
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis
                )
            }
        } else {
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
}

@Composable
fun SetForTileButton(
    onClick: Unit,
    modifier: Modifier = Modifier,
) {
    val context = LocalContext.current
    val setSuccessMessage = stringResource(R.string.set_for_tile_success)

    Button(
        onClick = {
            onClick
            Toast.makeText(
                context,
                setSuccessMessage,
                Toast.LENGTH_SHORT
            ).show()
        },
        colors = ButtonDefaults.buttonColors(
            backgroundColor = MaterialTheme.colors.secondary,
            contentColor = MaterialTheme.colors.onPrimary,
        ),
        modifier = modifier
    ) {
        Text(
            text = stringResource(R.string.set_for_tile),
        )
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
