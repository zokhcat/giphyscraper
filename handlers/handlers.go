package handlers

import (
	"giphyscraper/db"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func CreateAPIKey(c *gin.Context) {
	key := generateRandomKey(16)

	apiKey := db.APIKey{Key: key}
	result := db.A_DB.Create(&apiKey)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create API key"})
		return
	}

	c.JSON(201, gin.H{"id": apiKey.ID, "key": apiKey.Key})
}

func DeleteAPIKey(c *gin.Context) {
	if db.A_DB == nil {
		c.JSON(500, gin.H{"error": "Database not initialized"})
		return
	}

	id := c.Param("id")

	var apiKey db.APIKey
	result := db.A_DB.First(&apiKey, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "API key not found"})
		return
	}

	db.A_DB.Delete(&apiKey)

	c.JSON(200, gin.H{"message": "API key deleted"})
}

func generateRandomKey(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(uint64(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
