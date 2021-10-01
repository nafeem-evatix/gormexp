package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getSQLiteGormDBMust(filename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec("PRAGMA foreign_keys=ON;")
	return db
}
