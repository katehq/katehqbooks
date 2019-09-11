package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func OpenDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if nil != err {
		panic(err)
	}
	return db
}
