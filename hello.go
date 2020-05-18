package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	name string
	age  int
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.POST("/hello/:name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": "adf",
			"nick":    "adadf",
		})
	})

	router.Run(":3000")
}
