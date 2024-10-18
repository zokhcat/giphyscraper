package middleware

import (
	"giphyscraper/db"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("API-Key")
		if apiKey == "" {
			ctx.JSON(401, gin.H{"error": "API key is missing"})
			ctx.Abort()
			return
		}

		var apiKeyRecord db.APIKey
		result := db.A_DB.First(&apiKeyRecord, "key =?", apiKey)
		if result.Error != nil {
			ctx.JSON(401, gin.H{"error": "Invalid API Key"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
