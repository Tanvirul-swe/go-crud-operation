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
	r.POST("/"+constants.ApiVersion+"/create-post", controllers.PostCreate)
	r.GET("/"+constants.ApiVersion+"/posts", controllers.PostList)
	r.GET("/"+constants.ApiVersion+"/post/:Id", controllers.GetSinglePostById)
	r.PUT("/"+constants.ApiVersion+"/update-post/:Id", controllers.PostUpdate)
	r.DELETE("/"+constants.ApiVersion+"/delete-post/:Id", controllers.PostDelete)
	r.POST("/"+constants.ApiVersion+"/create-user", controllers.CreateUser)
	r.POST("/"+constants.ApiVersion+"/user-login", controllers.UserLogin)
	r.GET("/"+constants.ApiVersion+"/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
	// listen and serve on
}
