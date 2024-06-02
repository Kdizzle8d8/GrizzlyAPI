package main

import (
	"net/http"
	"strconv"

	"github.com/Kdizzle8d8/GrizzlyAPI/calendars"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", test)
	e.GET("/date/:month/:day", func(c echo.Context) error {
		month, _ := strconv.Atoi(c.Param("month"))
		day, _ := strconv.Atoi(c.Param("day"))

		res, err := calendars.GetTypeFromDate(month, day)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, res)
	})

	e.Start(":8080")
}

func test(c echo.Context) error {
	response := map[string]string{
		"message": "Hello World",
	}
	return c.JSON(http.StatusOK, response)
}
