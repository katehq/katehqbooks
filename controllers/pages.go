package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
)

func AddPage(c *gin.Context)  {
	c.HTML(200, "addpage.html", gin.H{

	})
}

func AddPagePOST(c *gin.Context)  {
	title := c.PostForm("title")
	content := c.PostForm("content")
	page := models.Page{
		Title: title,
		Content: content,
	}
	page.InsertPage()
}