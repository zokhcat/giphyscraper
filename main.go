package main

import (
	"giphyscraper/db"
	"giphyscraper/handlers"
	"giphyscraper/scrape_gif"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.POST("/api-key", handlers.CreateAPIKey)
	r.DELETE("/api-key/:id", handlers.DeleteAPIKey)
	r.Run()
	scrape_gif.Scrape()
}
