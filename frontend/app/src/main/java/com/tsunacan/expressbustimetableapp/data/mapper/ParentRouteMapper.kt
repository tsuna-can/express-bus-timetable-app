package com.tsunacan.expressbustimetableapp.data.mapper

import com.tsunacan.expressbustimetableapp.data.model.ParentRouteApiModel
import com.tsunacan.expressbustimetableapp.models.ParentRoute

object ParentRouteMapper {
    fun mapToParentRouteList(parentRouteList: List<ParentRouteApiModel>): List<ParentRoute> {
        return parentRouteList.map { parentRouteApiModel ->
            ParentRoute(
                parentRouteId = parentRouteApiModel.parentRouteId,
                parentRouteName = parentRouteApiModel.parentRouteName
            )
        }
    }
}
