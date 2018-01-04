package main

import (
	"os"

	"github.com/ivancduran/api-example/api"
	"github.com/ivancduran/api-example/models"
	"github.com/ivancduran/api-example/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	e.GET("/", api.IndexAPI)
	e.GET("/robots.txt", api.RobotsAPI)
	e.GET("/favicon.ico", api.FaviconAPI)
	e.GET("/ping", api.PingAPI)

	// v1
	a := e.Group("/v1")
	routes.ApiRoutes(a)

	// db init
	StartDB()

	// start server
	e.Logger.Fatal(e.Start(":8000"))
}

func StartDB() {
	models.AutoMigrate()
}
