package routes

import "github.com/gin-gonic/gin"

func UsersRoutes(r *gin.RouterGroup) {

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": nil,
		})
	})
}