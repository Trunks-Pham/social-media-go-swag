package main

import (
	"log"
	"social-media-backend/controllers"
	"social-media-backend/database"

	_ "social-media-backend/docs" // Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Social Media API
// @version 1.0
// @description This is a sample social media API.

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	database.ConnectDatabase()

	r := gin.Default()

	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Post routes
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
