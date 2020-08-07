package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func testResponse(c echo.Context) error {
	response := map[string]string{"ping": "pong"}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.GET("/test", testResponse)
	e.Logger.Fatal(e.Start(":8020"))
}
