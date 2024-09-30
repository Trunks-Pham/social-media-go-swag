package controllers

import (
	"net/http"
	"social-media-backend/database"
	"social-media-backend/models"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with provided data
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input", Error: err.Error()})
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create user", Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a single user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found", Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update user with provided data by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "Updated user data"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found", Error: err.Error()})
		return
	}
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input", Error: err.Error()})
		return
	}
	database.DB.Model(&user).Updates(updatedUser)
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a single user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found", Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Message: "User deleted successfully"})
}
