package com.tsunacan.expressbustimetableapp.presentation.ui.parentroutelist

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
import com.tsunacan.expressbustimetableapp.models.ParentRoute

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun ParentRouteListScreen(
    navigationToBusStopList: (String) -> Unit,
    modifier: Modifier = Modifier,
    viewModel: ParentRouteListScreenViewModel = hiltViewModel(),
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()
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
            when (uiState) {
                is ParentRouteListScreenUiState.Loaded -> {
                    val parentRouteList =
                        (uiState as ParentRouteListScreenUiState.Loaded).parentRouteList
                    parentRouteList.forEach { parentRoute ->
                        item {
                            ParentRouteChip(
                                parentRoute = parentRoute,
                                navigationToParentRoute = navigationToBusStopList,
                                modifier = contentModifier
                            )
                        }
                    }
                }

                ParentRouteListScreenUiState.Loading -> {
                    item {}
                }

                ParentRouteListScreenUiState.Failed -> {
                    item {}
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
