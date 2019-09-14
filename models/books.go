package models

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"katehqbooks/utils"
)
// Book model 
type Book struct {
	gorm.Model
	Title  string
	Author string
	Link   string
	Desc   string
	Cover  string
}
// InsertBook a book instance to a db
func (book *Book) InsertBook() {
	db := utils.OpenDB()
	defer db.Close()
	db.Create(book)
}
// GetAll get all of the books from the db
func (book *Book) GetAll() []Book {
	db := utils.OpenDB()
	defer db.Close()
	var books []Book
	db.Find(&books)
	return books
}
// FindByID get a instance by id
func (book *Book) FindByID(id int) {
	db := utils.OpenDB()
	defer db.Close()
	db.Where("id = ?", id).Find(&book)
}
// GetRandom get books random limit 5
func (book *Book) GetRandom() []Book {
	db := utils.OpenDB()
	defer db.Close()
	var books []Book
	db.Order(gorm.Expr("random()")).Limit(5).Find(&books)
	return books
}

// GetLatest get latest 8 books
func (book *Book) GetLatest() []Book {
	var books []Book
	db := utils.OpenDB()
	defer db.Close()
	db.Order("created_at desc").Limit(8).Find(&books)
	bookdetailsub(books)
	return books
}

func bookdetailsub(books []Book){
	for idx := range books {
		books[idx].Desc = utils.Substring(books[idx].Desc, 0,200) + "..."
	}
}