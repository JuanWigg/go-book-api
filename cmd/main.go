package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var lastId = 3
var books = []Book{
	{ID: "1", Title: "The Subtle Art of Not Giving a F*ck", Author: "Mark Manson", Description: "A book about life"},
	{ID: "2", Title: "The Holy Bible", Author: "Juan", Description: "Is the bible bro"},
	{ID: "3", Title: "Sandokan", Author: " Emilio Salgari", Description: "Libro sobre las aventuras del pirata Sandokan"},
	{ID: "4", Title: "Hola desde la versión 0.0.3", Author: "Juan Wiggenhauser", Description: "Test"},
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/ping", Ping)
	api.GET("/books", getBooks)
	api.GET("/book/:id", getBookById)
	api.POST("/book", addBook)
	api.PATCH("/book", updateBook)
	api.DELETE("/book/:id", deleteBookById)

	router.Run(":80")
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Book not found",
	})
}

func addBook(c *gin.Context) {
	var book Book
	c.BindJSON(&book)
	lastId += 1
	book.ID = strconv.Itoa(lastId)
	books = append(books, book)
	c.JSON(http.StatusOK, gin.H{
		"message": "Book added",
	})
}

func updateBook(c *gin.Context) {
	var bookToUpdate Book
	c.BindJSON(&bookToUpdate)
	for index, book := range books {
		if book.ID == bookToUpdate.ID {
			books[index].Author = bookToUpdate.Author
			books[index].Title = bookToUpdate.Title
			books[index].Description = bookToUpdate.Description
			c.JSON(http.StatusOK, gin.H{
				"message": "Book updated",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Book not found",
	})
}

func deleteBookById(c *gin.Context) {
	id := c.Param("id")
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Book deleted",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Book not found",
	})
}
