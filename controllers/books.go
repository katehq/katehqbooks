package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
)

// AddBook handle the get request from the `/books/new`
func AddBook(c *gin.Context) {
	c.HTML(200, "addbook.html", gin.H{})
}

// AddBookPOST handle the post request from the `books/new`
func AddBookPOST(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	link := c.PostForm("link")
	desc := c.PostForm("desc")
	cover := c.PostForm("cover")
	book := models.Book{
		Title:  title,
		Author: author,
		Link:   link,
		Desc:   desc,
		Cover:  cover,
	}
	book.InsertBook()
}
