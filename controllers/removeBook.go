package controllers

import (
	"fmt"
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

type NoBook struct {
	ISBN       int    `json:"isbn"`
	NumberBook uint64 `json:"numberbook"`
}

func RemoveBook(c *gin.Context) {
	var nb NoBook
	if err := c.BindJSON(&nb); err != nil {
		return
	}
	b := models.BookInventory{}

	b.ISBN = nb.ISBN

	check := b.CheckBook()

	if !check {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book is not available"})
		return
	} else {
		// _, err := b.DecreaseBook(nb.NumberBook)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		n := nb.NumberBook
		fmt.Println(n)
		if err := models.DB.Where("ISBN = ?", b.ISBN).Find(&b).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
			return
		}
		if b.TotalCopies == uint64(0) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book is not available"})
			return
		}
		if b.AvailableCopies == uint64(0) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "All Books issued"})
			return
		}
		if b.AvailableCopies < n {
			n = n - (n - b.AvailableCopies)
		}
		b.TotalCopies = b.TotalCopies - n
		b.AvailableCopies = b.AvailableCopies - n

		if err := models.DB.Where("ISBN = ?", b.ISBN).Save(&b).Error; err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book removed"})
}
