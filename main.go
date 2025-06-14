package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Specify icons with /icon/:name")
	})

	router.GET("/icon/:name", func(c *gin.Context) {
		iconName := c.Param("name")
		iconPath := filepath.Join("icons", iconName+".svg")

		c.File(iconPath)
	})

	router.Run(":3001") // listens on localhost:3001
}
