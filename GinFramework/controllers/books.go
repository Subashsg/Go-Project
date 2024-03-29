package controllers

import (
	"GinFramework/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

type CreateBookInput struct {
	Title       string `json:"title" binding: "required"`
	Author      string `json:"author" binding: "required"`
	Publication string `json:"Publication" binding: "required"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author, Publication: input.Publication}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {

	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"Publication"`
}

func updateBook(c *gin.Context) {

	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input UpdateBookInput

	err = c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&book).Update(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {

	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
