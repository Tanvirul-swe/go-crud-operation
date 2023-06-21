package main

import (
	"log"

	"example.com/go-crud/database"
	model "example.com/go-crud/models"
)

func main() {
	// Here we are calling the Init() function from initializers\database.go
	// This function will load the environment variables and connect to the database
	database.Init()

	// Here we are calling the AutoMigrate() function from gorm
	// This function will create the table in the database
	// If the table already exists then it will not create the table
	// It will only create the table if the table does not exist
	// If you want to drop the table and create the table again then you can use DropTable() function
	// But be careful while using this function because it will drop the table and all the data will be lost
	// So use this function only when you need to drop the table and create the table again
	// If you want to add a new column to the table then you can use AutoMigrate() function
	err :=
	database.DB.AutoMigrate(&model.Books{})
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.Category{})
	database.DB.AutoMigrate(&model.Followers{})
	database.DB.AutoMigrate(&model.Comments{})
	database.DB.AutoMigrate(&model.UserFevoriteBook{})
	database.DB.AutoMigrate(&model.Rating{})
	if err != nil {
		log.Fatal("Error while creating table")
	}

}
