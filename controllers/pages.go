package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
	"katehqbooks/utils"
)

// AddPage handle the get request from `/pages/new` mainly display the form
func AddPage(c *gin.Context) {
	c.HTML(200, "addpage.html", gin.H{})
}

// AddPagePOST handle the post request from the `/pages/new`
func AddPagePOST(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	page := models.Page{
		Title:   title,
		Content: content,
	}
	page.InsertPage()
}

// ShowPage show the detail of the page
func ShowPage(c *gin.Context) {
	titleParam := c.Param("title")
	var page models.Page
	page.GetPage(titleParam)
	title := page.Title
	content := page.Content
	safetext := utils.SafeHTML(content)
	c.HTML(200, "pagedetail.html", gin.H{
		"title":   title,
		"content": safetext,
	})
}
