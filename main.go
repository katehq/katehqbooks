package main

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/controllers"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("views/*")
	//GET
	r.GET("/", controllers.Index)
	r.GET("/migrate", controllers.Migrate)
	r.GET("/books/new", controllers.AddBook)
	r.GET("/books/show/:id", controllers.BookDetail)
	r.GET("/pages/new", controllers.AddPage)
	r.GET("/page/:title", controllers.ShowPage)
	//POST
	r.POST("/books/new", controllers.AddBookPOST)

	r.Run(":8080")
}
