package controllers

import (
	"net/http"
	"social-media-backend/database"
	"social-media-backend/models"

	"github.com/gin-gonic/gin"
)

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with given title and content
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Post true "Post data"
// @Success 200 {object} models.Post
// @Failure 400 {object} ErrorResponse
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input", Error: err.Error()})
		return
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create post", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPost godoc
// @Summary Get a post by ID
// @Description Get a single post by its ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {object} ErrorResponse
// @Router /posts/{id} [get]
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Post not found", Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update a post by ID
// @Description Update post with new title and content by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body models.Post true "Updated post data"
// @Success 200 {object} models.Post
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /posts/{id} [put]
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Post not found", Error: err.Error()})
		return
	}

	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input", Error: err.Error()})
		return
	}

	if err := database.DB.Model(&post).Updates(updatedPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to update post", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost godoc
// @Summary Delete a post by ID
// @Description Delete a single post by its ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.Delete(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Post not found", Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Message: "Post deleted successfully"})
}
