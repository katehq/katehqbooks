package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"katehqbooks/utils"
)

// Page model
type Page struct {
	gorm.Model
	Title   string
	Content string
}

// GetPage get page by title
func (p *Page) GetPage(title string) {
	db := utils.OpenDB()
	defer db.Close()
	db.Where("title=?", title).Find(&p)
}

// InsertPage insert a page instance
func (p *Page) InsertPage() {
	db := utils.OpenDB()
	defer db.Close()
	db.Create(p)
}
