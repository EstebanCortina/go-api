package controllers

import "github.com/gin-gonic/gin"

func UsersControllers(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "users",
		"data":  nil,
	})
}
