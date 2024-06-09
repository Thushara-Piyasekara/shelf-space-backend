package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/authors", getAuthors)
	router.Run("localhost:8080")
}
