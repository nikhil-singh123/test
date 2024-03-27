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

	// postBody1, _ := json.Marshal(b.ISBN)
	// postBody2, _ := json.Marshal(b.Title)
	// body = " ISBN Number : " + string(postBody1) + "      Book Title : " + string(postBody2)
	// subject := "New book is added"

	check := b.CheckBook()

	if !check {
		_, err := b.NewBook()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// go SendMsg(subject, body)

	} else {
		_, err := b.IncreaseBook()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// go SendMsg(subject, body)

		c.JSON(http.StatusAccepted, gin.H{"message": "Book is already existed. Increased the Total no. of copies of book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "query passed"})
}
