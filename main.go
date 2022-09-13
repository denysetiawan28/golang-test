package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-test/config"
	"net/http"
	"os"
)

func main() {
	// Echo instance
	e := echo.New()
	if _, err := os.Stat("./resources"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Printf("error : %s", err)
	}

	config.ConfigApps("./resources")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/path/:pathId", checkPath)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func checkPath(c echo.Context) error {
	par := c.Request().Header.Get("pathId")
	str := "exists"
	if _, err := os.Stat(par); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Printf("error : %s", err)
		str = err.Error()
	}
	return c.String(http.StatusOK, str)
}