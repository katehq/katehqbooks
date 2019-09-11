package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"katehqbooks/utils"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	Link   string
	Desc   string
}

func (book *Book) InsertBook() {
	db := utils.OpenDB()
	defer db.Close()
	db.Create(book)
}

func (book *Book) GetAll() []Book {
	db := utils.OpenDB()
	defer db.Close()
	var books []Book
	db.Find(&books)
	return books
}