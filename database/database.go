package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	dbname := "camapigns_db"
	username := "root"
	password := "yuleyek"
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic("Can't connect to the database")
	}
	return db
}
