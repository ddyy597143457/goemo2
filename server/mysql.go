package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func MysqlInit() {
	var err error
	db, err = gorm.Open("mysql", "root:@/emo?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func GetDBEngine() *gorm.DB {
	return db
}
