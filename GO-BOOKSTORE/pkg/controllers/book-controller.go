package controllers

import (
	"GO-BOOKSTORE/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	c.JSON(http.StatusOK, newBooks)
}

func GetBookID(c *gin.Context) {
	bookID := c.Param("bookid")
	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	bookDetails, _ := models.GetBookID(ID)
	c.JSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {
	createBook := &models.Book{}
	c.ShouldBind(createBook)
	b := createBook.CreateBook()
	c.JSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("bookid")
	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	del := models.DeleteBook(ID)
	//if err != nil {
	//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete details"})
	//return
	//}
	c.JSON(http.StatusOK, del)
}

func UpdateBook(c *gin.Context) {
	updateBook := &models.Book{}
	c.ShouldBindJSON(updateBook)
	bookID := c.Param("bookid")
	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing"})
		return
	}
	bookDetails, db := models.GetBookID(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	c.JSON(http.StatusOK, bookDetails)
}
