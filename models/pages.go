package models

import (
	"github.com/jinzhu/gorm"
	"katehqbooks/utils"
)

type Page struct {
	gorm.Model
	Title   string
	Content string
}

func (p *Page) GetPage(title string) {
	db := utils.OpenDB()
	defer db.Close()
	db.Where("title=?", title).Find(&p)
}

func (p *Page) InsertPage() {
	db := utils.OpenDB()
	defer db.Close()
	db.Create(p)
}
