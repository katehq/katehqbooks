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
	Cover  string
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

func (book *Book) FindByID(id int) {
	db := utils.OpenDB()
	defer db.Close()
	db.Where("id = ?", id).Find(&book)
}

func (book *Book) GetRandom() []Book {
	db := utils.OpenDB()
	defer db.Close()
	var books []Book
	db.Order(gorm.Expr("random()")).Limit(5).Find(&books)
	return books
}

func (book *Book) GetLatest() []Book {
	var books []Book
	db := utils.OpenDB()
	defer db.Close()
	db.Order("created_at desc").Limit(8).Find(&books)
	return books
}