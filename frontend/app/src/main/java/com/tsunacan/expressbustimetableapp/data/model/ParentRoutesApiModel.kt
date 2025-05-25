package com.tsunacan.expressbustimetableapp.data.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class ParentRoutesApiModel(
    @SerialName("parent_routes")
    val parentRoutes: List<ParentRouteApiModel>
)

@Serializable
data class ParentRouteApiModel(
    @SerialName("parent_route_id")
    val parentRouteId: String,
    @SerialName("parent_route_name")
    val parentRouteName: String,
)
