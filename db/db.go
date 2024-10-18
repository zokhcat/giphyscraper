package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var A_DB *gorm.DB

type APIKey struct {
	gorm.Model
	Key string `gorm:"unique;not null"`
}

func InitDB() {
	api_db, err := gorm.Open(sqlite.Open("api_keys.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	A_DB = api_db

	A_DB.AutoMigrate(&APIKey{})
}
