package services

import (
	"net/http"

	"example.com/go-crud/database"
	model "example.com/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostList(c *gin.Context) {

	//Get all posts
	var posts []model.Post
	database.DB.Find(&posts)

	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     "Posts List",
		"status_code": 200,
		"posts":       posts,
	})

}
