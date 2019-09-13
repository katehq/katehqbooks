package models

import (
	"katehqbooks/utils"
	"github.com/jinzhu/gorm"
)

type Page struct {
	gorm.Model
	Title string
	Content string
}

func (p *Page) GetPage(title string) {
	db := utils.OpenDB()
	defer db.Close()
	db.Where("title=?", title).Find(&p)
}

func (p *Page) InsertPage()  {
	db := utils.OpenDB()
	defer db.Close()
	db.Create(p)
}