package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ISBN            int    `json:"isbn"`
	LibID           int    `json:"libid"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     uint64 `json:"totalcopies"`
	AvailableCopies uint64 `json:"availablecopies"`
}

func AddBook(c *gin.Context) {
	var book Book
	// var body string

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "json err"})
		return
	}
	b := models.BookInventory{}

	b.ISBN = book.ISBN
	b.LibID = book.LibID
	b.Title = book.Title
	b.Author = book.Author
	b.Publisher = book.Publisher
	b.Version = book.Version
	b.TotalCopies = book.TotalCopies
	b.AvailableCopies = book.AvailableCopies

	check := b.CheckBook()

	if !check {
		_, err := b.NewBook()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	} else {
		_, err := b.IncreaseBook()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "Book already exists. Increased the Total no. of copies of book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "query passed"})
}
