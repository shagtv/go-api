package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var Db gorm.DB

func Connect() {
	var err error
	Db, err = gorm.Open("mysql", "root:000000@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	Db.SingularTable(true)
	Db.LogMode(true)
}