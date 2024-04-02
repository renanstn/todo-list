package main

import (
	"todo/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()
	database.SetupDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
