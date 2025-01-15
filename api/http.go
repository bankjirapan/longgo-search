package api

import (
	"github.com/gin-gonic/gin"
	"longgo-search.com/crawler"
	search "longgo-search.com/search"
)

func StartServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/crawler", func(c *gin.Context) {
		go crawler.Crawler()
		c.JSON(200, gin.H{
			"message": "Crawling runing...",
		})
	})

	r.GET("/search", func(c *gin.Context) {
		text := c.Query("t")

		result := search.Search(text)

		c.JSON(200, gin.H{
			"search": text,
			"match":  result,
		})
	})

	r.Run()
}
