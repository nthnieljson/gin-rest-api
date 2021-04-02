package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("get request on ping wwwwwwwwwwwwwwwww")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port) 
}