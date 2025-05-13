package com.tsunacan.expressbustimetableapp.presentation.ui.parentroutelist

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.google.android.horologist.annotations.ExperimentalHorologistApi
import com.tsunacan.expressbustimetableapp.data.repository.ParentRouteRepository
import com.tsunacan.expressbustimetableapp.models.ParentRoute
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.catch
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.onStart
import kotlinx.coroutines.flow.stateIn
import javax.inject.Inject

@ExperimentalHorologistApi
@HiltViewModel
class ParentRouteListScreenViewModel @Inject constructor(
    parentRouteRepository: ParentRouteRepository
) : ViewModel() {

    val uiState: StateFlow<ParentRouteListScreenUiState> = parentRouteRepository.getParentRouteList()
        .map<List<ParentRoute>, ParentRouteListScreenUiState> { parentRoutes ->
            ParentRouteListScreenUiState.Loaded(parentRoutes)
        }
        .onStart { emit(ParentRouteListScreenUiState.Loading) }
        .catch {
            emit(ParentRouteListScreenUiState.Failed)
        }
        .stateIn(viewModelScope, SharingStarted.Lazily, ParentRouteListScreenUiState.Loading)

}
