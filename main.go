package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/days", DaysHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func DaysHandler(c echo.Context) error {
	numberOfDays := strconv.Itoa(CountDays())

	return c.String(http.StatusOK, numberOfDays)
}

func CountDays() int {
	startTime := time.Now()
	endTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	timeToEnd := endTime.Sub(startTime)

	return int(timeToEnd.Hours()) / 24
}
