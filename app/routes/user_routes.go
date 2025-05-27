package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sf966gpm/GinLab/controllers"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/:id", controllers.GetUserByID)
	}
}
