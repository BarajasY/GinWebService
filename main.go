package main

import (
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.InitializeDB()
	

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Listening on port 8080
	r.Run()
}
