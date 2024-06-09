package routes

import (
	"shelf-space/rest-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	router.GET("/books", controllers.GetBooks)
	router.GET("/authors", controllers.GetAuthors)
	router.Run("localhost:8080")
}
