package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
	"strconv"
)

func BookDetail(c *gin.Context) {
	id_param := c.Param("id")
	id, err := strconv.Atoi(id_param)
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
	var books []models.Book
	books = book.GetRandom()
	c.HTML(200, "bookdetail.html", gin.H{
		"title":  title,
		"author": author,
		"link":   link,
		"desc":   desc,
		"cover":  cover,
		"books":  books,
	})
}
