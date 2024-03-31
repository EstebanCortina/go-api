package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.RouterGroup) {
	UsersRoutes(r.Group("/users"))
	PostsRoutes(r.Group("/posts"))
}
