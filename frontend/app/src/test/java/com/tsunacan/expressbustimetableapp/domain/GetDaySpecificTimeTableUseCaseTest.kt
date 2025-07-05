package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.models.Timetable
import com.tsunacan.expressbustimetableapp.models.TimetableEntry
import com.tsunacan.expressbustimetableapp.testdouble.FakeTimetableRepository
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import java.time.DayOfWeek
import java.time.LocalDate
import java.time.LocalTime

class GetDaySpecificTimetableUseCaseTest {

    private lateinit var fakeTimetableRepository: FakeTimetableRepository
    private lateinit var getDaySpecificTimetableUseCase: GetDaySpecificTimetableUseCase

    @Before
    fun setUp() {
        fakeTimetableRepository = FakeTimetableRepository()
        getDaySpecificTimetableUseCase = GetDaySpecificTimetableUseCase(fakeTimetableRepository)
    }

    @Test
    fun `invoke filters and sorts timetable by day of week and departure time`() = runBlocking {
        // Arrange
        val parentRouteId = "route1"
        val busStopId = "stop1"
        val today = LocalDate.of(2023, 10, 10) // Tuesday

        val timetableEntries = listOf(
            TimetableEntry(
                departureTime = LocalTime.of(10, 0),
                destination = "Destination A",
                availableDayOfWeek = setOf(DayOfWeek.MONDAY, DayOfWeek.TUESDAY)
            ),
            TimetableEntry(
                departureTime = LocalTime.of(9, 0),
                destination = "Destination B",
                availableDayOfWeek = setOf(DayOfWeek.TUESDAY)
            ),
            TimetableEntry(
                departureTime = LocalTime.of(11, 0),
                destination = "Destination C",
                availableDayOfWeek = setOf(DayOfWeek.WEDNESDAY)
            )
        )

        val timetable = Timetable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timetableEntryList = timetableEntries
        )

        fakeTimetableRepository.addTimetable(parentRouteId, busStopId, timetable)

        // Act
        val result = getDaySpecificTimetableUseCase(parentRouteId, busStopId, today)

        // Assert
        val expectedEntries = listOf(
            timetableEntries[1], // 9:00, Tuesday
            timetableEntries[0], // 10:00, Tuesday
        )
        assertEquals(expectedEntries, result.timetableEntryList)
    }

    @Test
    fun `invoke returns empty list when no entries match the day of week`() = runBlocking {
        // Arrange
        val parentRouteId = "route1"
        val busStopId = "stop1"
        val today = LocalDate.of(2023, 10, 11) // Wednesday

        val timetableEntries = listOf(
            TimetableEntry(
                departureTime = LocalTime.of(10, 0),
                destination = "Destination A",
                availableDayOfWeek = setOf(DayOfWeek.MONDAY, DayOfWeek.TUESDAY)
            )
        )

        val timetable = Timetable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timetableEntryList = timetableEntries
        )

        fakeTimetableRepository.addTimetable(parentRouteId, busStopId, timetable)

        // Act
        val result = getDaySpecificTimetableUseCase(parentRouteId, busStopId, today)

        // Assert
        assertEquals(emptyList<TimetableEntry>(), result.timetableEntryList)
    }
}
