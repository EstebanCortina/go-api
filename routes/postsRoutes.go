package routes

import (
	"crud/controllers"

	"github.com/gin-gonic/gin"
)

func PostsRoutes(r *gin.RouterGroup) {
	r.GET("", controllers.PostsController)
}
