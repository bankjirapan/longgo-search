package api

import (
	"github.com/gin-gonic/gin"
	search "longgo-search.com/search"
)

func StartServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/search", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		text := c.Query("t")

		result := search.Search(text)

		c.JSON(200, gin.H{
			"search": text,
			"match":  result,
		})
	})

	r.Run(":8090")
}
