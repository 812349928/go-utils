package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbUtil = new(gorm.DB)

func InitDb(connArgs string) {
	DbUtil, _ = gorm.Open("mysql", connArgs)

	//析构函数
	//runtime.SetFinalizer(db, UnQuery)
}

func UnQuery(db *gorm.DB) {
	db.Close()
}
