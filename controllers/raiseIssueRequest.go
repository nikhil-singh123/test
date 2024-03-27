package controllers

import (
	"net/http"
	"test/models"
	"time"

	"github.com/gin-gonic/gin"
)

type RaiseIssue struct {
	BookID int `json:"bookid"`
}

func RaiseIssueRequest(c *gin.Context) {
	Email := c.Param("email")
	var ri RaiseIssue
	if err := c.BindJSON(&ri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "json err"})
		return
	}
	var book models.BookInventory
	if err := models.DB.Where("isbn = ?", ri.BookID).Find(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book is not Available"})
		return
	}
	if book.TotalCopies == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book is not Available"})
		return
	}
	if book.AvailableCopies <= 0 && book.TotalCopies > uint64(0) {
		c.JSON(http.StatusNotFound, gin.H{"message": "All Books are already Issued"})
		return
	}
	var user models.Users
	if err := models.DB.Where("email = ?", Email).Find(&user).Error; err != nil {
		return
	}

	var requestEvents models.RequestEvents
	requestEvents.BookID = ri.BookID
	requestEvents.ReaderId = user.ID
	requestEvents.RequestDate = time.Now()
	requestEvents.ApprovelDate = nil
	requestEvents.RequestType = "issue"
	if err := models.DB.Create(&requestEvents).Error; err != nil {
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Raise Issue Request"})

}
