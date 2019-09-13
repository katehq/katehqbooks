package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
)

func AddPage(c *gin.Context) {
	c.HTML(200, "addpage.html", gin.H{})
}

func AddPagePOST(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	page := models.Page{
		Title:   title,
		Content: content,
	}
	page.InsertPage()
}

func ShowPage(c *gin.Context) {
	title_param := c.Param("title")
	var page models.Page
	page.GetPage(title_param)
	title := page.Title
	content := page.Content
	c.HTML(200, "pagedetail.html", gin.H{
		"title":   title,
		"content": content,
	})
}
