@file:OptIn(ExperimentalHorologistApi::class)

package com.tsunacan.expressbustimetableapp.tile

import android.content.Context
import androidx.wear.protolayout.ActionBuilders
import androidx.wear.protolayout.ColorBuilders.argb
import androidx.wear.protolayout.DeviceParametersBuilders.DeviceParameters
import androidx.wear.protolayout.LayoutElementBuilders
import androidx.wear.protolayout.ModifiersBuilders
import androidx.wear.protolayout.ResourceBuilders.Resources
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
import com.tsunacan.expressbustimetableapp.models.Trip

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

private fun mainTileLayout(
    context: Context,
    deviceParameters: DeviceParameters,
    state: MainTileState
) = PrimaryLayout.Builder(deviceParameters)
    .setResponsiveContentInsetEnabled(true)
    .setPrimaryLabelTextContent(
        Text.Builder(context, context.getString(R.string.from) + state.stopName)
            .setColor(argb(Colors.DEFAULT.onSurface))
            .setTypography(androidx.wear.protolayout.material.Typography.TYPOGRAPHY_CAPTION1)
            .build()
    )
    .setContent(
        LayoutElementBuilders.Column.Builder()
            .apply {
                state.timeTable.take(3).forEach { trip ->
                    addContent(
                        Text.Builder(context, trip.arrivalTime + " " + trip.destination)
                            .setColor(argb(Colors.DEFAULT.onSurface))
                            .setTypography(androidx.wear.protolayout.material.Typography.TYPOGRAPHY_CAPTION3)
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
    return TilePreviewData() { request ->
        MainTileRenderer(context).renderTimeline(
            MainTileState(
                "Tokyo",
                listOf(
                    Trip("test", "12:00"),
                    Trip("test", "12:30"),
                    Trip("test", "13:00"),
                    Trip("test", "13:30"),
                )
            ),
            request
        )
    }
}
