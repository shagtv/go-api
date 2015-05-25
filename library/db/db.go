package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MysqlConnect *gorm.DB

func Connection() *gorm.DB {
	if MysqlConnect == nil {
		connect, err := gorm.Open("mysql", "root:000000@/test?charset=utf8&parseTime=True&loc=Local")

		if err != nil {
			panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
		}
		connect.SingularTable(true)
		connect.LogMode(true)

		MysqlConnect = &connect
	}

	return MysqlConnect
}
