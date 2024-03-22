package routes

import (
	"GO-BOOKSTORE/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterBookStore = func(router *gin.Engine) {

	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/:bookid", controllers.GetBookID)
	router.PUT("/book//:bookid", controllers.UpdateBook)
	router.DELETE("/book//:bookid", controllers.DeleteBook)
}
