package main

import (
	"crud/config"
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvironment()
}

func main() {
	app := gin.Default()

	api := app.Group("/api")

	{
		routes.ApiRoutes(api)
	}

	app.Run()
}
