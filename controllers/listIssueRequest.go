package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func ListIssueRequest(c *gin.Context) {
	var requestEvents []models.RequestEvents
	issue := "issue"
	if err := models.DB.Where("request_type = ?", issue).Find(&requestEvents).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No issue request"})
		return
	}
	c.JSON(http.StatusOK, requestEvents)
}
