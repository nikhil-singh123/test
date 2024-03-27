package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Name string `json:"name" gorm:"unique"`
}

func RegLibrary(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "input err"})
		return
	}
	l := models.Library{}

	l.Name = input.Name

	check := l.CheckName()

	if !check {
		_, err := l.SaveLibrary()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name already exists. Choose another name"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "query passed"})

}
