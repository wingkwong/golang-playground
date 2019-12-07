package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/test", endpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/test", endpoint)
	}

	router.Run(":8080")
}

func endpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello-world",
	})
}
