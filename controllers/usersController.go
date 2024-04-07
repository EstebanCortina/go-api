package controllers

import (
	"crud/config"
	"crud/handlers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func UsersControllers(c *gin.Context) {
	out := make(chan models.User)
	var results []models.User

	go handlers.ExecQuery(
		config.ConnectDB(),
		"SELECT * FROM user", out)

	for r := range out {
		results = append(results, r)
	}

	c.JSON(200, gin.H{
		"title": "users",
		"data":  results,
	})
}
