package controllers

import (
	"example.com/go-crud/services"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {

  services.CreateCategory(c)
}

func CategoriesList(c *gin.Context) {
	
  services.GetCategories(c)
}