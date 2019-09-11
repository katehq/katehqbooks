package models

import (
	"katehqbooks/utils"
)

func AutoMigrate() {
	db := utils.OpenDB()
	defer db.Close()
	db.AutoMigrate(&Book{})
}
