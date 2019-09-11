package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
)

func AddBook(c *gin.Context) {
	c.HTML(200, "addbook.html", gin.H{})
}

func AddBookPOST(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	link := c.PostForm("link")
	desc := c.PostForm("desc")
	book := models.Book{
		Title:  title,
		Author: author,
		Link:   link,
		Desc:   desc,
	}
	book.InsertBook()
}
