@file:OptIn(ExperimentalHorologistApi::class)

package com.tsunacan.expressbustimetableapp.tile

import android.content.Context
import androidx.wear.protolayout.ActionBuilders
import androidx.wear.protolayout.ColorBuilders.argb
import androidx.wear.protolayout.DeviceParametersBuilders.DeviceParameters
import androidx.wear.protolayout.DimensionBuilders.dp
import androidx.wear.protolayout.LayoutElementBuilders
import androidx.wear.protolayout.ModifiersBuilders
import androidx.wear.protolayout.ResourceBuilders.Resources
import androidx.wear.protolayout.material.Typography
import androidx.wear.protolayout.material.Colors
import androidx.wear.protolayout.material.CompactChip
import androidx.wear.protolayout.material.Text
import androidx.wear.protolayout.material.layouts.PrimaryLayout
import androidx.wear.tiles.tooling.preview.Preview
import androidx.wear.tiles.tooling.preview.TilePreviewData
import androidx.wear.tooling.preview.devices.WearDevices
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.google.android.horologist.tiles.render.SingleTileLayoutRenderer
import com.tsunacan.expressbustimetableapp.R
import com.tsunacan.expressbustimetableapp.models.DepartureTimeAndDestination
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter

class MainTileRenderer(context: Context) :
    SingleTileLayoutRenderer<MainTileState, Unit>(context) {

    override fun renderTile(
        state: MainTileState,
        deviceParameters: DeviceParameters
    ): LayoutElementBuilders.LayoutElement {
        return mainTileLayout(context, deviceParameters, state)
    }

    override fun Resources.Builder.produceRequestedResources(
        resourceState: Unit,
        deviceParameters: DeviceParameters,
        resourceIds: List<String>
    ) {
        // No resources to produce
    }
}

private val formatter = DateTimeFormatter.ofPattern("HH:mm")

private fun mainTileLayout(
    context: Context,
    deviceParameters: DeviceParameters,
    state: MainTileState
) = PrimaryLayout.Builder(deviceParameters)
    .setResponsiveContentInsetEnabled(true)
    .setPrimaryLabelTextContent(
        Text.Builder(context, state.parentRouteName)
            .setColor(argb(Colors.DEFAULT.onSurface))
            .setTypography(Typography.TYPOGRAPHY_CAPTION1)
            .build()
    )
    .setContent(
        LayoutElementBuilders.Column.Builder()
            .addContent(
                Text.Builder(context, state.stopName)
                    .setColor(argb(Colors.DEFAULT.onSurface))
                    .setTypography(Typography.TYPOGRAPHY_CAPTION2)
                    .build()
            )
            .addContent(
                LayoutElementBuilders.Spacer.Builder()
                    .setHeight(dp(8f))
                    .build()
            )
            .apply {
                state.timeTable.take(3).forEach {
                    val formattedTime = it.departureTime.format(formatter)
                    addContent(
                        Text.Builder(context, formattedTime + " " + it.destination)
                            .setColor(argb(Colors.DEFAULT.onSurface))
                            .setTypography(Typography.TYPOGRAPHY_CAPTION1)
                            .build()
                    )
                }
            }
            .build()
    )
    .setPrimaryChipContent(
        CompactChip.Builder(
            context,
            context.getString(R.string.more_info),
            emptyClickable,
            deviceParameters
        ).build()
    ).build()

val emptyClickable = ModifiersBuilders.Clickable.Builder()
    .setOnClick(ActionBuilders.LoadAction.Builder().build())
    .setId("")
    .build()

@Preview(device = WearDevices.SMALL_ROUND)
@Preview(device = WearDevices.LARGE_ROUND)
fun mainTileLayoutPreview(context: Context): TilePreviewData {
    val dummyTime1 = LocalDateTime.now().plusMinutes(10).toLocalTime()
    val dummyTime2 = LocalDateTime.now().plusMinutes(20).toLocalTime()
    val dummyTime3 = LocalDateTime.now().plusMinutes(30).toLocalTime()
    return TilePreviewData() { request ->
        MainTileRenderer(context).renderTimeline(
            MainTileState(
                parentRouteName = "Nagoya-go",
                stopName = "Tokyo",
                listOf(
                    DepartureTimeAndDestination(dummyTime1, "Tokyo"),
                    DepartureTimeAndDestination(dummyTime2, "Sapporo"),
                    DepartureTimeAndDestination(dummyTime3, "Chiba"),
                )
            ),
            request
        )
    }
}
