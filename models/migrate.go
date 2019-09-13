package models

import (
	"katehqbooks/utils"
)
// AutoMigrate migrate db
func AutoMigrate() {
	db := utils.OpenDB()
	defer db.Close()
	db.AutoMigrate(&Book{}, &Page{})
}
