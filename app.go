package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 10},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 20},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 30},
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, books)
	})
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, book := range books {
			if book.ID == id {
				c.IndentedJSON(http.StatusOK, book)
				return
			}
		}
		c.AbortWithError(http.StatusNotFound, errors.New("Book not found"))
	})
	router.POST("/", func(c *gin.Context) {
		var book book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		books = append(books, book)
		c.IndentedJSON(http.StatusOK, book)
	})
	router.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, book := range books {
			if book.ID == id {
				book := book
				if err := c.ShouldBindJSON(&book); err != nil {
					c.AbortWithError(http.StatusBadRequest, err)
					return
				}
				books[i] = book
				c.IndentedJSON(http.StatusOK, book)
				return
			}
		}
		c.AbortWithError(http.StatusNotFound, errors.New("Book not found"))
	})

	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, book := range books {
			if book.ID == id {
				books = append(books[:i], books[i+1:]...)
				c.IndentedJSON(http.StatusOK, book)
				return
			}
		}
		c.AbortWithError(http.StatusNotFound, errors.New("Book not found"))
	})
	router.Run("localhost:1234")
}
