package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
	"strconv"
	"katehqbooks/utils"
)

// BookDetail display the detail page
func BookDetail(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic("sorry")
	}
	var book models.Book
	book.FindByID(id)
	title := book.Title
	author := book.Author
	cover := book.Cover
	link := book.Link
	desc := book.Desc
	safetext := utils.SafeHTML(desc)
	var books []models.Book
	books = book.GetRandom()
	c.HTML(200, "bookdetail.html", gin.H{
		"title":  title,
		"author": author,
		"link":   link,
		"desc":   safetext,
		"cover":  cover,
		"books":  books,
	})
}
