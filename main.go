package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/books", listBookHandler)

	router.POST("/books", createBookHandler)

	router.DELETE("/books/:id", deleteBookHandler)

	router.Run(":3333")
}

type Book struct {
	ID 			string `json:"id"`
	Title		string `json:"title"`
	Author  string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Game of Thrones", Author: "George R. R. Martin"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func listBookHandler(context *gin.Context) {
	context.JSON(http.StatusOK, books)
}

func createBookHandler(context *gin.Context) {
	var book Book
	if err := context.ShouldBindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, book)

	context.JSON(http.StatusCreated, book)
}

func deleteBookHandler(context *gin.Context) {
	id := context.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	context.Status(http.StatusNoContent)
}