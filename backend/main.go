package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// DB接続
func init() {
	dsn := "host=my-postgres-container user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
}

// NetworkTimeTableとDepartureTimeAndDestinationApiModel構造体を定義
type NetworkTimeTable struct {
	ParentRouteId      string                   `json:"parentRouteId"`
	ParentRouteName    string                   `json:"parentRouteName"`
	StopId             string                   `json:"stopId"`
	StopName           string                   `json:"stopName"`
	TimeTableEntryList []TimeTableEntryApiModel `json:"timeTableEntryList"`
}

type TimeTableEntryApiModel struct {
	DepartureTime      string `json:"departureTime"`
	Destination        string `json:"destination"`
	AvailableDayOfWeek []int  `json:"availableDayOfWeek"`
}

// BusStopApiModel構造体を定義
type BusStopApiModel struct {
	ParentRouteId   string `json:"parentRouteId"`
	ParentRouteName string `json:"parentRouteName"`
	StopId          string `json:"stopId"`
	StopName        string `json:"stopName"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/bus-stops", getBusStops)
	e.GET("/timetable", getTimeTable)

	e.Logger.Fatal(e.Start(":8080"))
}

func getTimeTable(c echo.Context) error {
	parentRouteId := c.QueryParam("parent-route-id")
	busStopId := c.QueryParam("bus-stop-id")

	response := NetworkTimeTable{
		ParentRouteId:   parentRouteId,
		ParentRouteName: parentRouteId,
		StopId:          busStopId,
		StopName:        busStopId,
		TimeTableEntryList: []TimeTableEntryApiModel{
			{DepartureTime: "08:00", Destination: "City Center", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "10:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "12:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "14:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "16:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "18:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "19:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "20:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "20:15", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "20:17", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "21:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "22:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "23:00", Destination: "Airport", AvailableDayOfWeek: []int{1, 3}},
			{DepartureTime: "23:30", Destination: "Airport", AvailableDayOfWeek: []int{1, 2, 3, 4, 5, 6, 7}},
		},
	}
	// レスポンスを標準出力に出力
	printResponse(response)

	return c.JSON(http.StatusOK, response)
}

func getBusStops(c echo.Context) error {
	response := []BusStopApiModel{
		{ParentRouteId: "route-1", ParentRouteName: "Route 1", StopId: "stop-1", StopName: "Station A"},
		{ParentRouteId: "route-2", ParentRouteName: "Route 2", StopId: "stop-2", StopName: "Station B"},
		{ParentRouteId: "route-3", ParentRouteName: "Route 3", StopId: "stop-3", StopName: "Station C"},
	}

	return c.JSON(http.StatusOK, response)
}

// レスポンスを標準出力に出力するユーティリティ関数
func printResponse(response interface{}) {
	// JSON文字列に変換
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal response: %v\n", err)
		return
	}
	// 標準出力に表示
	fmt.Println(string(jsonData))
}
