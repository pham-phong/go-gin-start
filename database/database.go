package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "go_module"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
