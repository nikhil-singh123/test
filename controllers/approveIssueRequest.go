package controllers

import (
	"net/http"
	"test/models"
	"time"

	"github.com/gin-gonic/gin"
)

type DetailReqID struct {
	ReqID int `json:"reqid"`
}

func ApproveRejectIssueRequest(c *gin.Context) {
	Email := c.Param("email")

	var ri DetailReqID
	if err := c.BindJSON(&ri); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var user models.Users
	if err := models.DB.Where("email = ?", Email).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := user.ID

	var requestEvents models.RequestEvents
	if err := models.DB.Where("req_id = ?", ri.ReqID).Find(&requestEvents).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Request ID Not Found"})
		return
	}
	if requestEvents.RequestType == "issued" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book is already approved"})
		return
	}

	var book models.BookInventory
	if err := models.DB.Where("isbn = ?", requestEvents.BookID).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Book is not available "})
		return
	}
	if book.TotalCopies < uint64(1) {
		c.JSON(http.StatusAccepted, gin.H{"message": "Book is not available "})
		return
	}
	if book.AvailableCopies <= uint64(0) && book.TotalCopies >= uint64(1) {
		c.JSON(http.StatusAccepted, gin.H{"message": " All books are already issued"})
		return
	}
	book.AvailableCopies = book.AvailableCopies - 1
	if err := models.DB.Where("isbn = ?", requestEvents.BookID).Save(&book).Error; err != nil {
		c.JSON(http.StatusAccepted, gin.H{"message": "error in updating book inventory"})
		return
	}

	t := time.Now()
	requestEvents.ApprovelDate = &t
	requestEvents.ApproverID = user.ID
	requestEvents.RequestType = "issued"

	if err := models.DB.Where("req_id = ?", ri.ReqID).Save(&requestEvents).Error; err != nil {
		c.JSON(http.StatusAccepted, gin.H{"message": "request event not saved"})
		return
	}

	var issueRegistery models.IssueRegistery

	issueRegistery.ISBN = requestEvents.BookID
	issueRegistery.ReaderId = requestEvents.ReaderId
	issueRegistery.IssueApproverID = id
	issueRegistery.IssueStatus = "approved"
	issueRegistery.IssueDate = time.Now()
	issueRegistery.ExpectedReturnDate = time.Now().AddDate(0, 0, 7)

	if err := models.DB.Create(&issueRegistery).Error; err != nil {
		c.JSON(http.StatusAccepted, gin.H{"message": "issue registry not saved"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book approved"})

}
