package controllers

import (
	"net/http"

	"example.com/go-crud/database"
	model "example.com/go-crud/models"
	"example.com/go-crud/services"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//Get model values from request
	var post model.Post
	//Bind JSON body to model
	c.BindJSON(&post)
	// Check if title and body fields are not empty
	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Title Field is required",
			"status_code": http.StatusBadRequest,
		})
		return
	}
	// Check if title and body fields are not empty

	if post.Body == "" {
		c.JSON(400, gin.H{
			"message":     "Body Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	//Create Post
	result := database.DB.Create(&post)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Post Creation Failed",
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	//Return response as JSON with status code 200
	c.JSON(http.StatusCreated, gin.H{
		"message":     "Post Created Successfully",
		"status_code": http.StatusCreated,
		"post":        post,
	})

}

func PostList(c *gin.Context) {
	// Call the service to get the posts
	services.PostList(c)

}

func GetSinglePostById(c *gin.Context) {

	//Get post id from request
	postId := c.Param("Id")

	//Get post
	var post model.Post
	result := database.DB.First(&post, postId)
	// Return result as JSON response with status code 400 if there is an error or post not found in database
	// RowsAffected is 0 if no record found

	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Post Not Found",
		})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     "Post Details",
		"status_code": http.StatusOK,
		"post":        post,
	})

}

func PostUpdate(c *gin.Context) {

	//Get post id from request
	postId := c.Param("Id")

	//Get model values from request
	var post model.Post
	c.BindJSON(&post)

	//Bind JSON body to model
	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Title Field is required",
			"status_code": http.StatusBadRequest,
		})
		return
	}
	// Check if title and body fields are not empty

	if post.Body == "" {
		c.JSON(400, gin.H{
			"message":     "Body Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	//Update Post
	result := database.DB.Model(&post).Where("id = ?", postId).Updates(&post)

	// Return result as JSON response with status code 400 if there is an error or post not found in database
	// RowsAffected is 0 if no record found

	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Post Not Found",
		})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	//Return response as JSON with status code 200
	c.JSON(http.StatusAccepted, gin.H{
		"message":     "Post Updated Successfully",
		"status_code": http.StatusAccepted,
		"post":        post,
	})

}

func PostDelete(c *gin.Context) {

	//Get post id from request
	postId := c.Param("Id")

	//Get post
	var post model.Post
	result := database.DB.First(&post, postId)
	// Return result as JSON response with status code 400 if there is an error or post not found in database
	// RowsAffected is 0 if no record found

	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Post Not Found",
		})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	//Delete Post
	database.DB.Delete(&post)

	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     "Post Deleted Successfully",
		"status_code": http.StatusAccepted,
	})
}
