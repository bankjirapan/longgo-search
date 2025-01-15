package api

import (
	"github.com/gin-gonic/gin"
	"longgo-search.com/crawler"
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

	r.Run()
}
