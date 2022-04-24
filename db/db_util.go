package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbUtil struct {
	Db *gorm.DB
}

func (d *DbUtil) Init(connArgs string) {
	d.Db, _ = gorm.Open("mysql", connArgs)

	//析构函数
	//runtime.SetFinalizer(db, UnQuery)
}

func UnQuery(db *gorm.DB) {
	db.Close()
}
