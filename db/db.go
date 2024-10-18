package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type APIKey struct {
	ID  uint   `gorm:"primaryKey;default:auto_random()"`
	Key string `gorm:"unique;not null"`
}

func InitDB() *gorm.DB {
	api_db, err := gorm.Open(sqlite.Open("api_keys.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	api_db.AutoMigrate(&APIKey{})

	return api_db
}

func CreateAPIKey(api_db *gorm.DB, apikey APIKey) error {
	result := api_db.Create(&apikey)
	return result.Error
}

func DeleteAPIKey(api_db *gorm.DB, userID uint) error {
	result := api_db.Where("user_id = ?", userID).Delete(&APIKey{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
