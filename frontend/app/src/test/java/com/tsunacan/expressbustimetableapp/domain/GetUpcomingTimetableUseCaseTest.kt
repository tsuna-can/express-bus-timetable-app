package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.models.Timetable
import com.tsunacan.expressbustimetableapp.models.TimetableEntry
import io.mockk.coEvery
import io.mockk.every
import io.mockk.mockk
import io.mockk.mockkStatic
import io.mockk.unmockkStatic
import kotlinx.coroutines.runBlocking
import org.junit.After
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import java.time.DayOfWeek
import java.time.LocalDate
import java.time.LocalDateTime
import java.time.LocalTime
import java.time.format.DateTimeFormatter

class GetUpcomingTimetableUseCaseTest {
    private lateinit var getUpcomingTimetableUseCase: GetUpcomingTimetableUseCase
    private lateinit var getDaySpecificTimetableUseCase: GetDaySpecificTimetableUseCase

    // Test constants
    private val parentRouteId = "route1"
    private val busStopId = "stop1"
    private val today = LocalDate.of(2023, 5, 15)
    private val fixedCurrentTime = LocalTime.of(12, 0) // Noon

    @Before
    fun setup() {
        getDaySpecificTimetableUseCase = mockk()
        getUpcomingTimetableUseCase = GetUpcomingTimetableUseCase(getDaySpecificTimetableUseCase)

        // Mock the LocalDateTime.now() static method
        mockkStatic(LocalDateTime::class)
        val mockedNow = LocalDateTime.of(today, fixedCurrentTime)
        every { LocalDateTime.now() } returns mockedNow
    }

    @After
    fun tearDown() {
        // Clean up the static mock after tests
        unmockkStatic(LocalDateTime::class)
    }

    private fun createTimetableEntry(
        time: String,
        destination: String = "Test Destination"
    ): TimetableEntry {
        return TimetableEntry(
            departureTime = LocalTime.parse(time, DateTimeFormatter.ofPattern("HH:mm")),
            destination = destination,
            availableDayOfWeek = setOf(DayOfWeek.MONDAY, DayOfWeek.TUESDAY, DayOfWeek.WEDNESDAY)
        )
    }

    @Test
    fun `when there are upcoming departures, should return up to 3 entries`() = runBlocking {
        // Given
        val mockTimetable = Timetable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timetableEntryList =
            listOf(
                createTimetableEntry("10:00"),
                createTimetableEntry("11:00"),
                createTimetableEntry("12:30"), // First upcoming
                createTimetableEntry("13:00"), // Second upcoming
                createTimetableEntry("14:00"), // Third upcoming
                createTimetableEntry("15:00") // Not included
            )
        )

        coEvery {
            getDaySpecificTimetableUseCase.invoke(
                parentRouteId,
                busStopId,
                today
            )
        } returns mockTimetable

        // When
        val result = getUpcomingTimetableUseCase(parentRouteId, busStopId, today)

        // Then
        assertEquals(3, result.timetableEntryList.size)
        assertEquals(LocalTime.parse("12:30"), result.timetableEntryList[0].departureTime)
        assertEquals(LocalTime.parse("13:00"), result.timetableEntryList[1].departureTime)
        assertEquals(LocalTime.parse("14:00"), result.timetableEntryList[2].departureTime)
    }

    @Test
    fun `when there are no upcoming departures, should return empty list`() = runBlocking {
        // Given
        val mockTimetable = Timetable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timetableEntryList =
            listOf(
                createTimetableEntry("10:00"),
                createTimetableEntry("11:00")
                // All departures are before current time (12:00)
            )
        )

        coEvery {
            getDaySpecificTimetableUseCase.invoke(
                parentRouteId,
                busStopId,
                today
            )
        } returns mockTimetable

        // When
        val result = getUpcomingTimetableUseCase(parentRouteId, busStopId, today)

        // Then
        assertEquals(0, result.timetableEntryList.size)
    }

    @Test
    fun `when there are less than 3 upcoming departures, should return all available`() =
        runBlocking {
            // Given
            val mockTimetable = Timetable(
                parentRouteId = parentRouteId,
                parentRouteName = "Route 1",
                stopId = busStopId,
                stopName = "Stop 1",
                timetableEntryList =
                listOf(
                    createTimetableEntry("10:00"),
                    createTimetableEntry("11:00"),
                    createTimetableEntry("12:30"), // First upcoming
                    createTimetableEntry("13:00") // Second upcoming
                )
            )

            coEvery {
                getDaySpecificTimetableUseCase.invoke(
                    parentRouteId,
                    busStopId,
                    today
                )
            } returns mockTimetable

            // When
            val result = getUpcomingTimetableUseCase(parentRouteId, busStopId, today)

            // Then
            assertEquals(2, result.timetableEntryList.size)
            assertEquals(LocalTime.parse("12:30"), result.timetableEntryList[0].departureTime)
            assertEquals(LocalTime.parse("13:00"), result.timetableEntryList[1].departureTime)
        }

    @Test
    fun `when the first departure is at the current time, it should be included`() = runBlocking {
        // Given
        val mockTimetable = Timetable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timetableEntryList =
            listOf(
                createTimetableEntry("10:00"),
                createTimetableEntry("12:00"), // Equal to current time, should be included
                createTimetableEntry("12:30"),
                createTimetableEntry("13:00")
            )
        )

        coEvery {
            getDaySpecificTimetableUseCase.invoke(
                parentRouteId,
                busStopId,
                today
            )
        } returns mockTimetable

        // When
        val result = getUpcomingTimetableUseCase(parentRouteId, busStopId, today)

        // Then
        assertEquals(3, result.timetableEntryList.size)
        assertEquals(LocalTime.parse("12:00"), result.timetableEntryList[0].departureTime)
        assertEquals(LocalTime.parse("12:30"), result.timetableEntryList[1].departureTime)
        assertEquals(LocalTime.parse("13:00"), result.timetableEntryList[2].departureTime)
    }
}
