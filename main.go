package main

import (
	"crud/config"
	"crud/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic!:", r)
		}
	}()
	config.LoadEnvironment()
	config.ConnectDB()
}

func main() {
	gin.ForceConsoleColor()
	app := gin.Default()

	api := app.Group("/api")

	{
		routes.ApiRoutes(api)
	}

	app.Run()
}
