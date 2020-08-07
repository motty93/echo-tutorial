package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User is struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users   = map[int]*User{}
	userSeq = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &User{
		ID: userSeq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[user.ID] = u
	userSeq++

	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name

	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)

	return c.NoContent(http.StatusNoContent)
}

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

	// users
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8020"))
}
