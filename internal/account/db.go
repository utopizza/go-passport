package account

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = createSQLite()
}

func createSQLite() *gorm.DB {
	// dir and file
	dbDir := "./data"
	dbFile := path.Join(dbDir, "sqlite.db")

	// create database
	os.MkdirAll(dbDir, 0755)
	os.Create(dbFile)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	// create tables
	tables := []string{
		"account_info.sql",
		"account_connect.sql",
		"account_device.sql",
		"account_email.sql",
		"account_mobile.sql",
	}
	for _, tb := range tables {
		filePath := path.Join("../../configs/schema", tb)
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(string(bytes))
		if err != nil {
			panic(err)
		}
	}

	// return gorm connect
	gormDb, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return gormDb
}

func GetDB() *gorm.DB {
	return db
}
