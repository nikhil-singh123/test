package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Query string `json:"query"`
}

func SearchBook(c *gin.Context) {
	var query Query

	query.Query = c.Param("query")

	if len(query.Query) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Query is empty, write 'title/author/publisher' of book"})
		return
	}
	// var b models.BookInventory
	b, err := models.FindBook(models.Query(query))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}
	if len(b) != 0 {
		c.JSON(http.StatusOK, b)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
	}

}
