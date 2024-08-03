package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Kdizzle8d8/GrizzlyAPI/calendars"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", test)
	e.GET("/date/:month/:day", func(c echo.Context) error {
		fmt.Println("Hello World")
		month, _ := strconv.Atoi(c.Param("month"))
		day, _ := strconv.Atoi(c.Param("day"))

		res, err := calendars.GetTypeFromDate(month, day)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, res)
	})
	e.GET("/until/first", func(c echo.Context) error {
		firstDay := time.Date(2024, time.August, 14, 0, 0, 0, 0, time.Local)
		now := time.Now()
		fmt.Println(firstDay, now)
		duration := firstDay.Sub(now)

		days := int(duration.Hours() / 24)
		hours := int(duration.Hours()) % 24
		minutes := int(duration.Minutes()) % 60

		response := map[string]interface{}{
			"days":    days,
			"hours":   hours,
			"minutes": minutes,
		}

		return c.JSON(http.StatusOK, response)
	})

	e.Start(":8080")
}

func test(c echo.Context) error {
	response := map[string]string{
		"message": "Hello World",
	}
	return c.JSON(http.StatusOK, response)
}
