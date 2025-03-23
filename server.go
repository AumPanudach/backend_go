package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"users": []string{"Alice", "Bob"}})
	})
	router.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"items": []string{"Item1", "Item2"}})
	})
	router.Run(":8080")
}
