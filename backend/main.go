package main

import (
  "github.com/tsuna-can/express-bus-time-table-app/backend/infrastructure"
)

// @title Express Bus Time Table API
// @version 1.0
// @description This is a sample server for Express Bus Time Table API.
// @contact.name API Support
// @contact.url http://www.example.com/support
func main() {
	infrastructure.InitRouter()
}

