package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getAuthors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, authors)
}
