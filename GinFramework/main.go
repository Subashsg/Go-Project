package main

import (
	"GinFramework/controllers"
	"GinFramework/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/books", controllers.FindBooks)
	r.POST("/api/books", controllers.CreateBook)
	r.GET("/api/books/:id", controllers.FindBook)
	r.DELETE("/api/books/:id", controllers.DeleteBook)

	r.Run()
}
