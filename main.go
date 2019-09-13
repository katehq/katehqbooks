package main

import (
	"os"
	"katehqbooks/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	r.POST("/pages/new", controllers.AddPagePOST)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	r.Run(PORT)
}
