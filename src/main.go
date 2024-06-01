package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main(){
	e:=echo.New()
	e.GET("/",test)

	e.Start(":8080")
}

func test(c echo.Context)error{
	response := map[string]string{
		"message":"Hello World",
	}
	return c.JSON(http.StatusOK,response)
}