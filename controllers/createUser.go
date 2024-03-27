package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

type RegUsers struct {
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"not null;unique" validate:"required"`
	ContactNumber int    `json:"contactnumber"`
	Role          string `json:"role"`
	LibID         int    `json:"libid"`
}

func UserCreate(c *gin.Context) {
	var input RegUsers
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bind json"})
		return
	}
	u := models.Users{}

	u.Name = input.Name
	u.Email = input.Email
	u.ContactNumber = input.ContactNumber
	u.Role = input.Role
	u.LibID = input.LibID

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "query passed"})

}
