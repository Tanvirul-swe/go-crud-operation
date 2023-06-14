package controllers

import (
	"example.com/go-crud/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Call the service to get the posts
	services.CreateUser(c)

}

func UserLogin(c *gin.Context) {
	// Call the service to get the posts
	services.UserLogin(c)

}

func Validate(c *gin.Context) {
	// Call the service to get the posts
	services.Validate(c)
}
