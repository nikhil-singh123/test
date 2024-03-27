package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func UpdateBook(c *gin.Context) {
	var ub models.BookInventory

	if err := c.BindJSON(&ub); err != nil {
		return
	}

	check := ub.CheckBook()

	if !check {
		_, err := ub.NewBook()
		if err != nil {
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "No book found. Added this new book to the inventory"})
		return
	} else {
		_, err := ub.UpdateDetails()
		if err != nil {
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "Book Details Updated"})
		return
	}

}

func UpdateISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	var book models.BookInventory
	if err := models.DB.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		return
	}
	c.JSON(http.StatusOK, book)
}
