// routes/routes.go
package routes

import (
	"social-media-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("/", controllers.CreatePost)
		postRoutes.GET("/:id", controllers.GetPost)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}
}
