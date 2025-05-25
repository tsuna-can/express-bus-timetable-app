package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.models.TimeTable
import com.tsunacan.expressbustimetableapp.models.TimeTableEntry
import com.tsunacan.expressbustimetableapp.testdouble.FakeTimetableRepository
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import java.time.DayOfWeek
import java.time.LocalDate
import java.time.LocalTime

class GetDaySpecificTimeTableUseCaseTest {

    private lateinit var fakeTimetableRepository: FakeTimetableRepository
    private lateinit var getDaySpecificTimeTableUseCase: GetDaySpecificTimeTableUseCase

    @Before
    fun setUp() {
        fakeTimetableRepository = FakeTimetableRepository()
        getDaySpecificTimeTableUseCase = GetDaySpecificTimeTableUseCase(fakeTimetableRepository)
    }

    @Test
    fun `invoke filters and sorts timetable by day of week and departure time`() = runBlocking {
        // Arrange
        val parentRouteId = "route1"
        val busStopId = "stop1"
        val today = LocalDate.of(2023, 10, 10) // Tuesday

        val timeTableEntries = listOf(
            TimeTableEntry(
                departureTime = LocalTime.of(10, 0),
                destination = "Destination A",
                availableDayOfWeek = setOf(DayOfWeek.MONDAY, DayOfWeek.TUESDAY)
            ),
            TimeTableEntry(
                departureTime = LocalTime.of(9, 0),
                destination = "Destination B",
                availableDayOfWeek = setOf(DayOfWeek.TUESDAY)
            ),
            TimeTableEntry(
                departureTime = LocalTime.of(11, 0),
                destination = "Destination C",
                availableDayOfWeek = setOf(DayOfWeek.WEDNESDAY)
            )
        )

        val timeTable = TimeTable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timeTableEntryList = timeTableEntries
        )

        fakeTimetableRepository.addTimeTable(parentRouteId, busStopId, timeTable)

        // Act
        val result = getDaySpecificTimeTableUseCase(parentRouteId, busStopId, today)

        // Assert
        val expectedEntries = listOf(
            timeTableEntries[1], // 9:00, Tuesday
            timeTableEntries[0], // 10:00, Tuesday
        )
        assertEquals(expectedEntries, result.timeTableEntryList)
    }

    @Test
    fun `invoke returns empty list when no entries match the day of week`() = runBlocking {
        // Arrange
        val parentRouteId = "route1"
        val busStopId = "stop1"
        val today = LocalDate.of(2023, 10, 11) // Wednesday

        val timeTableEntries = listOf(
            TimeTableEntry(
                departureTime = LocalTime.of(10, 0),
                destination = "Destination A",
                availableDayOfWeek = setOf(DayOfWeek.MONDAY, DayOfWeek.TUESDAY)
            )
        )

        val timeTable = TimeTable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timeTableEntryList = timeTableEntries
        )

        fakeTimetableRepository.addTimeTable(parentRouteId, busStopId, timeTable)

        // Act
        val result = getDaySpecificTimeTableUseCase(parentRouteId, busStopId, today)

        // Assert
        assertEquals(emptyList<TimeTableEntry>(), result.timeTableEntryList)
    }
}
