package controllers

import (
	"fmt"
	"net/http"

	"example.com/go-crud/database"
	model "example.com/go-crud/models"
	"example.com/go-crud/services"
	"github.com/gin-gonic/gin"
)

func BookCreate(c *gin.Context) {

	//Get model values from request
	var book model.Books
	//Bind JSON body to model
	c.BindJSON(&book)
	// Check if title and body fields are not empty
	fmt.Println("Author Name is : ", book.Author_Name)
	if book.Author_Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "AuthorName Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if book.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Name Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if book.Category == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Category Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	if book.Discription == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Discription Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if book.Book_Url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "BookUrl Field is required",
			"status_code": http.StatusBadRequest,
		})
		return
	}
	if book.Image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Image Field is required",
			"status_code": http.StatusBadRequest,
		})
		return
	}

	if book.User_Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "UserId Field is required",
			"status_code": http.StatusBadRequest,
		})
		return
	}

	if book.Status == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Status Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	//Create Post
	result := database.DB.Create(&book)

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
		"message":     "Created Successfully",
		"status_code": http.StatusCreated,
		"post":        book,
	})

}

func BooksList(c *gin.Context) {
	// Call the service to get the posts
	services.BooksList(c)

}

func GetSingleBookById(c *gin.Context) {

	//Get post id from request
	bookId := c.Param("Id")

	//Get post
	var book model.Books
	// result := database.DB.First(&book, bookId)
	// Return result as JSON response with status code 400 if there is an error or post not found in database
	// RowsAffected is 0 if no record found
	result:=database.DB.Table("books").Joins("left join books on books.category = categories.id").Scan(&book)

   fmt.Println("Book Id : ",bookId)
   fmt.Println("response is : ",book)
	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     "Book Not Found",
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
		"message":     "Book Details",
		"status_code": http.StatusOK,
		"book":    book,
	})

}

func BookUpdate(c *gin.Context) {

	//Get post id from request
	bookId := c.Param("Id")

	//Get model values from request
	var book model.Books
	//Bind JSON body to model

	c.BindJSON(&book)

	if book.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Name Field is required",
			"status_code": http.StatusBadRequest,
		})
		return
	}
	// Check if title and body fields are not empty

	if book.Author_Name == "" {
		c.JSON(400, gin.H{
			"message":     "AuthorName Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	//Update Post
	result := database.DB.Model(&book).Where("id = ?", bookId).Updates(&book)

	// Return result as JSON response with status code 400 if there is an error or post not found in database
	// RowsAffected is 0 if no record found

	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     "Book Not Found",
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
		"book":        book,
	})

}

func PostDelete(c *gin.Context) {

	//Get post id from request
	postId := c.Param("Id")

	//Get post
	var post model.Books
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
