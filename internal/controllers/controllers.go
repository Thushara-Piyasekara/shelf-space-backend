package controllers

import (
	"net/http"
	"shelf-space/rest-api/internal/constants"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, constants.Books)
}

func GetAuthors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, constants.Authors)
}
