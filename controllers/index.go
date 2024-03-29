package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
)

// Index home of my website
func Index(c *gin.Context) {
	var book models.Book
	books := book.GetLatest()
	c.HTML(200, "index.html", gin.H{
		"books": books,
	})
}
