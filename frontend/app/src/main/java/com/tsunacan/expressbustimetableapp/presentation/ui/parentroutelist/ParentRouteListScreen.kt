package com.tsunacan.expressbustimetableapp.presentation.ui.parentroutelist

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
import com.google.android.horologist.compose.layout.ScreenScaffold
import com.google.android.horologist.compose.layout.rememberResponsiveColumnState
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults.ItemType
import com.tsunacan.expressbustimetableapp.R
import com.tsunacan.expressbustimetableapp.models.ParentRoute
import com.tsunacan.expressbustimetableapp.presentation.ui.common.ErrorScreen
import com.tsunacan.expressbustimetableapp.presentation.ui.common.LoadingIndicator

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun ParentRouteListScreen(
    navigationToBusStopList: (String) -> Unit,
    modifier: Modifier = Modifier,
    viewModel: ParentRouteListScreenViewModel = hiltViewModel(),
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()

    ScreenScaffold {
        val contentModifier = modifier
            .fillMaxWidth()
            .padding(bottom = 8.dp)

        when (uiState) {
            is ParentRouteListScreenUiState.Loaded -> {
                ParentRouteListScreen(
                    parentRouteList = (uiState as ParentRouteListScreenUiState.Loaded).parentRouteList,
                    navigationToBusStopList = navigationToBusStopList,
                    modifier = contentModifier,
                )
            }

            ParentRouteListScreenUiState.Loading -> {
                LoadingIndicator()
            }

            ParentRouteListScreenUiState.Failed -> {
                ErrorScreen()
            }
        }
    }
}


@OptIn(ExperimentalHorologistApi::class)
@Composable
fun ParentRouteListScreen(
    parentRouteList: List<ParentRoute>,
    navigationToBusStopList: (String) -> Unit,
    modifier: Modifier = Modifier,
) {
    val listState = rememberResponsiveColumnState(
        contentPadding = ScalingLazyColumnDefaults.padding(
            first = ItemType.Chip,
            last = ItemType.Chip,
        ),
    )

    ScalingLazyColumn(
        columnState = listState,
    ) {
        item {
            Text(
                text = stringResource(R.string.route_list),
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
        if (parentRouteList.isEmpty()) {
            item {
                Text(
                    text = stringResource(R.string.no_route),
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis
                )
            }
        } else {
            parentRouteList.forEach { parentRoute ->
                item {
                    ParentRouteChip(
                        parentRoute = parentRoute,
                        navigationToParentRoute = navigationToBusStopList,
                        modifier = modifier
                    )
                }
            }
        }
    }
}

@Composable
fun ParentRouteChip(
    parentRoute: ParentRoute,
    navigationToParentRoute: (String) -> Unit,
    modifier: Modifier = Modifier,
) {
    Chip(
        modifier = modifier,
        onClick = { navigationToParentRoute(parentRoute.parentRouteId) },
        label = {
            Text(
                text = parentRoute.parentRouteName,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        }
    )
}
