package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbWrite, dbRead *gorm.DB

func InitDB(dsnWrite, dsnRead string) {
	dbWrite, err := gorm.Open("mysql", dsnWrite)
	if err != nil {
		panic(err)
	}

	dbRead, err = gorm.Open("mysql", dsnRead)
	if err != nil {
		dbRead = dbWrite
	}
}
