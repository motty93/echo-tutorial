package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// hoge routing追加
func hogeHandler(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("pash"))
}

func testResponse(c echo.Context) error {
	response := map[string]string{"ping": "pong"}

	return c.JSON(http.StatusOK, response)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/test", testResponse)
	// get hoge
	e.GET("/hoge/:path", hogeHandler)
	e.Logger.Fatal(e.Start(":8020"))
}
