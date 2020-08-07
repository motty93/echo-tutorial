package main

import (
	"net/http"

	"github.com/labstack/echo"
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
	e := echo.New()

	e.GET("/test", testResponse)
	// get hoge追加
	e.GET("/hoge/:path", hogeHandler)
	e.Logger.Fatal(e.Start(":8020"))
}
