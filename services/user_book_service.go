package services

import (
	"net/http"
	"example.com/go-crud/database"
	model "example.com/go-crud/models"
	"github.com/gin-gonic/gin"
)

func BooksList(c *gin.Context) {

	//Get all posts
	var book []model.Books
	database.DB.Find(&book)
  
	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     "Books List",
		"status_code": 200,
		"book":        book,
	})

}
