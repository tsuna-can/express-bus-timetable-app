package com.tsunacan.expressbustimetableapp.presentation.ui.busstop

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.wear.compose.material.Button
import androidx.wear.compose.material.Text
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.compose.layout.ScalingLazyColumn
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults
import com.google.android.horologist.compose.layout.ScalingLazyColumnDefaults.ItemType
import com.google.android.horologist.compose.layout.ScreenScaffold
import com.google.android.horologist.compose.layout.rememberResponsiveColumnState
import com.google.android.horologist.compose.material.Chip

@OptIn(ExperimentalHorologistApi::class)
@Composable
fun BusStopScreen(modifier: Modifier = Modifier) {
    val listState = rememberResponsiveColumnState(
        contentPadding = ScalingLazyColumnDefaults.padding(
            first = ItemType.Text,
            last = ItemType.Chip,
        ),
    )

    ScreenScaffold(
        scrollState = listState,
    ) {
        val contentModifier = Modifier
            .fillMaxWidth()
            .padding(bottom = 8.dp)

        ScalingLazyColumn(
            columnState = listState,
        ) {
            item {
                Text(
                    text = "Tokyo Express",
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis
                )
            }
            item {
                Text(
                    text = "Tokyo station",
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis
                )
            }
            item {
                Button(
                    onClick = { /* ... */ },
                    modifier = contentModifier,
                ) {
                    Text(
                        text = "Set as default",
                    )
                }
            }
            item {
                Chip(
                    modifier = contentModifier,
                    onClick = { /* ... */ },
                    label = {
                        Text(
                            text = "16:00 Kyoto",
                            maxLines = 1,
                            overflow = TextOverflow.Ellipsis
                        )
                    }
                )
            }
            item {
                Chip(
                    modifier = contentModifier,
                    onClick = { /* ... */ },
                    label = {
                        Text(
                            text = "16:30 Sapporo",
                            maxLines = 1,
                            overflow = TextOverflow.Ellipsis
                        )
                    }
                )
            }
        }
    }
}
