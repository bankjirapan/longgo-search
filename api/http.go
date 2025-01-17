package api

import (
	"net/http"

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
	r.GET("/", func(ctx *gin.Context) {
		r.LoadHTMLFiles("web/index.html")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "LongGo Search - Online document search engine",
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
