package com.tsunacan.expressbustimetableapp.domain

import com.tsunacan.expressbustimetableapp.models.TimeTable
import com.tsunacan.expressbustimetableapp.models.TimeTableEntry
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
    private lateinit var getUpcomingTimeTableUseCase: GetUpcomingTimeTableUseCase
    private lateinit var getDaySpecificTimeTableUseCase: GetDaySpecificTimeTableUseCase

    // Test constants
    private val parentRouteId = "route1"
    private val busStopId = "stop1"
    private val today = LocalDate.of(2023, 5, 15)
    private val fixedCurrentTime = LocalTime.of(12, 0) // Noon

    @Before
    fun setup() {
        getDaySpecificTimeTableUseCase = mockk()
        getUpcomingTimeTableUseCase = GetUpcomingTimeTableUseCase(getDaySpecificTimeTableUseCase)

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

    private fun createTimeTableEntry(
        time: String,
        destination: String = "Test Destination"
    ): TimeTableEntry {
        return TimeTableEntry(
            departureTime = LocalTime.parse(time, DateTimeFormatter.ofPattern("HH:mm")),
            destination = destination,
            availableDayOfWeek = setOf(DayOfWeek.MONDAY, DayOfWeek.TUESDAY, DayOfWeek.WEDNESDAY)
        )
    }

    @Test
    fun `when there are upcoming departures, should return up to 3 entries`() = runBlocking {
        // Given
        val mockTimeTable = TimeTable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timeTableEntryList =
                listOf(
                    createTimeTableEntry("10:00"),
                    createTimeTableEntry("11:00"),
                    createTimeTableEntry("12:30"), // First upcoming
                    createTimeTableEntry("13:00"), // Second upcoming
                    createTimeTableEntry("14:00"), // Third upcoming
                    createTimeTableEntry("15:00") // Not included
                )
        )

        coEvery {
            getDaySpecificTimeTableUseCase.invoke(
                parentRouteId,
                busStopId,
                today
            )
        } returns mockTimeTable

        // When
        val result = getUpcomingTimeTableUseCase(parentRouteId, busStopId, today)

        // Then
        assertEquals(3, result.timeTableEntryList.size)
        assertEquals(LocalTime.parse("12:30"), result.timeTableEntryList[0].departureTime)
        assertEquals(LocalTime.parse("13:00"), result.timeTableEntryList[1].departureTime)
        assertEquals(LocalTime.parse("14:00"), result.timeTableEntryList[2].departureTime)
    }

    @Test
    fun `when there are no upcoming departures, should return empty list`() = runBlocking {
        // Given
        val mockTimeTable = TimeTable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timeTableEntryList =
                listOf(
                    createTimeTableEntry("10:00"),
                    createTimeTableEntry("11:00")
                    // All departures are before current time (12:00)
                )
        )

        coEvery {
            getDaySpecificTimeTableUseCase.invoke(
                parentRouteId,
                busStopId,
                today
            )
        } returns mockTimeTable

        // When
        val result = getUpcomingTimeTableUseCase(parentRouteId, busStopId, today)

        // Then
        assertEquals(0, result.timeTableEntryList.size)
    }

    @Test
    fun `when there are less than 3 upcoming departures, should return all available`() =
        runBlocking {
            // Given
            val mockTimeTable = TimeTable(
                parentRouteId = parentRouteId,
                parentRouteName = "Route 1",
                stopId = busStopId,
                stopName = "Stop 1",
                timeTableEntryList =
                    listOf(
                        createTimeTableEntry("10:00"),
                        createTimeTableEntry("11:00"),
                        createTimeTableEntry("12:30"), // First upcoming
                        createTimeTableEntry("13:00")  // Second upcoming
                    )
            )

            coEvery {
                getDaySpecificTimeTableUseCase.invoke(
                    parentRouteId,
                    busStopId,
                    today
                )
            } returns mockTimeTable

            // When
            val result = getUpcomingTimeTableUseCase(parentRouteId, busStopId, today)

            // Then
            assertEquals(2, result.timeTableEntryList.size)
            assertEquals(LocalTime.parse("12:30"), result.timeTableEntryList[0].departureTime)
            assertEquals(LocalTime.parse("13:00"), result.timeTableEntryList[1].departureTime)
        }

    @Test
    fun `when the first departure is at the current time, it should be included`() = runBlocking {
        // Given
        val mockTimeTable = TimeTable(
            parentRouteId = parentRouteId,
            parentRouteName = "Route 1",
            stopId = busStopId,
            stopName = "Stop 1",
            timeTableEntryList =
                listOf(
                    createTimeTableEntry("10:00"),
                    createTimeTableEntry("12:00"), // Equal to current time, should be included
                    createTimeTableEntry("12:30"),
                    createTimeTableEntry("13:00")
                )
        )

        coEvery {
            getDaySpecificTimeTableUseCase.invoke(
                parentRouteId,
                busStopId,
                today
            )
        } returns mockTimeTable

        // When
        val result = getUpcomingTimeTableUseCase(parentRouteId, busStopId, today)

        // Then
        assertEquals(3, result.timeTableEntryList.size)
        assertEquals(LocalTime.parse("12:00"), result.timeTableEntryList[0].departureTime)
        assertEquals(LocalTime.parse("12:30"), result.timeTableEntryList[1].departureTime)
        assertEquals(LocalTime.parse("13:00"), result.timeTableEntryList[2].departureTime)
    }
}
