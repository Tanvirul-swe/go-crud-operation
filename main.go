package main

import (
	"fmt"
	"example.com/go-crud/constants"
	"example.com/go-crud/controllers"
	"example.com/go-crud/database"
	"example.com/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API Version : " + constants.ApiVersion)
	database.Init()
	r := gin.Default()
	r.POST("/"+constants.ApiVersion+"/create-book", controllers.BookCreate)
	r.GET("/"+constants.ApiVersion+"/books", controllers.BooksList)
	r.GET("/"+constants.ApiVersion+"/book/:Id", controllers.GetSingleBookById)
	// r.PUT("/"+constants.ApiVersion+"/update-book/:Id", controllers.BookUpdate)
	r.POST("/"+constants.ApiVersion+"/create-category", controllers.CreateCategory)
	r.GET("/"+constants.ApiVersion+"/categories", controllers.CategoriesList)
	// r.DELETE("/"+constants.ApiVersion+"/delete-post/:Id", controllers.PostDelete)
	r.POST("/"+constants.ApiVersion+"/create-user", controllers.CreateUser)
	r.POST("/"+constants.ApiVersion+"/user-login", controllers.UserLogin)
	r.GET("/"+constants.ApiVersion+"/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
	// listen and serve on
}
