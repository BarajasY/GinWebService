package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDB() {

	var conn string = "host=localhost user=postgres password=puerta756859 dbname=homework port=1402"

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database0")
	}

	err = db.AutoMigrate(&Homework{})

	if err != nil {
		return
	}

	Db = db
}
