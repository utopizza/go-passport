package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var defaultDB *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/passport?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defaultDB = db
}

func GetDB() *gorm.DB {
	return defaultDB
}
