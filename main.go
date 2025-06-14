package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("public/*")
	gin.SetMode(gin.ReleaseMode)

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/icon", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/icon/:name", func(c *gin.Context) {
		iconName := c.Param("name")
		iconPath := filepath.Join("icons", iconName+".svg")

		c.File(iconPath)
	})

	router.Run(":3001") // listens on localhost:3001
}
