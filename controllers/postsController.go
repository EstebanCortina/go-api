package controllers

import "github.com/gin-gonic/gin"

func PostsController(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "posts",
		"data":  nil,
	})
}
