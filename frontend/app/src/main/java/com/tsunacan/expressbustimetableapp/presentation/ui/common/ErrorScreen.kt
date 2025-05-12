package com.tsunacan.expressbustimetableapp.presentation.ui.common

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.wear.compose.material.Text
import com.tsunacan.expressbustimetableapp.R

@Composable
fun ErrorScreen(
    modifier: Modifier = Modifier,
    errorMessage: String = "",
) {
    val title = if (errorMessage.isEmpty()) {
        stringResource(R.string.error_screen_title)
    } else {
        errorMessage
    }

    Box(
        modifier = modifier.fillMaxSize(),
        contentAlignment = Alignment.Center
    ) {
        Text(
            text = title,
            modifier = modifier
                .padding(16.dp),
        )
    }

    // Todo : implement retry button
}