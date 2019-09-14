package utils

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
	T "html/template"
)


//OpenDB open a database easily
func OpenDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if nil != err {
		panic(err)
	}
	return db
}

//Substring son set of the original str
func Substring(str string, start int, end int) string {

	rs := []rune(str)
	length := len(rs)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	return string(rs[start:end])
}

// SafeHTML return the original type of the text.
func SafeHTML(str string) T.HTML {
	sp := strings.Split(str, "\n")
	var HTML string
	for _, val := range sp {
		HTML= HTML+`<p>`+val+`</p>`

	}
	return T.HTML(HTML)
}