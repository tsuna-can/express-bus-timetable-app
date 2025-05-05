package com.tsunacan.expressbustimetableapp.data.mapper

import com.tsunacan.expressbustimetableapp.data.model.BusStopsApiModel
import com.tsunacan.expressbustimetableapp.models.BusStop

object BusStopMapper {
    fun map(busStopsApiModel: BusStopsApiModel): List<BusStop> {
        val busStopList = busStopsApiModel.busStops
        val parentRouteId = busStopsApiModel.parentRouteId
        val parentRouteName = busStopsApiModel.parentRouteName
        return busStopList.map { busStop->
            BusStop(
                parentRouteId = parentRouteId,
                parentRouteName = parentRouteName,
                stopId = busStop.busStopId,
                stopName = busStop.busStopName
            )
        }
    }
}
