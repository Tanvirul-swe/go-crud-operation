package services

import (
	"net/http"

	"example.com/go-crud/database"
	model "example.com/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {

	//Get model values from request
	var category model.Category
	//Bind JSON body to model
	c.BindJSON(&category)
	// Check if title and body fields are not empty
	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Name Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	//Create Category
	result := database.DB.Create(&category)
	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Creation Failed",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     "Category Created",
		"status_code": 200,
		"category":    category,
	})

}

// Get all the categories

func GetCategories(c *gin.Context) {
	
	var categories []model.Category
	//Get all the categories
	result := database.DB.Find(&categories)
	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Fetching Failed",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     "Categories Fetched",
		"status_code": 200,
		"categories":  categories,
	})

}
