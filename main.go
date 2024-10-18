package main

import (
	"giphyscraper/db"
	"giphyscraper/handlers"
	"giphyscraper/middleware"
	"giphyscraper/scrape_gif"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.POST("/api-key", handlers.CreateAPIKey)
	r.DELETE("/api-key/:id", handlers.DeleteAPIKey)

	protected := r.Group("/")
	protected.Use(middleware.APIKeyMiddleware())
	{
		protected.GET("/giphy", scrape_gif.Scrape)
	}

	r.Run()
}
