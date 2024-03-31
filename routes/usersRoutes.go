package routes

import (
	"crud/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.RouterGroup) {
	r.GET("", controllers.UsersControllers)
}
