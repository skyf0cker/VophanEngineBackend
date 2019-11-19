package model

import (
"fmt"
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {

	var err error

	dbType := "mysql"
	dbName := "vophanEngine"
	user := "root"
	password := "root"
	host := "127.0.0.1"

	//connect to the database
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	//Debug
	//db.LogMode(true)

	if err != nil {
		//log.Println(err)
	}

	//全局禁用表名复数
	db.SingularTable(true)
	db.Set("gorm:auto_preload", true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserChoice{})
}

func CloseDB() {
	defer db.Close()
}

func GetDb() *gorm.DB {
	return db
}
