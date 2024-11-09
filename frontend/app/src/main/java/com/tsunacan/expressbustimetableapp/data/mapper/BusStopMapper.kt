package com.tsunacan.expressbustimetableapp.data.mapper

import com.tsunacan.expressbustimetableapp.data.model.BusStopApiModel
import com.tsunacan.expressbustimetableapp.models.BusStop

object BusStopMapper {
    fun map(busStopList: List<BusStopApiModel>): List<BusStop> {
        return busStopList.map { busStopApiModel ->
            BusStop(
                parentRouteId = busStopApiModel.parentRouteId,
                parentRouteName = busStopApiModel.parentRouteName,
                stopId = busStopApiModel.stopId,
                stopName = busStopApiModel.stopName
            )
        }
    }
}
