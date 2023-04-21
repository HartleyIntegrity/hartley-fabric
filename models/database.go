package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{}, &Contract{})

	DB = db
}
