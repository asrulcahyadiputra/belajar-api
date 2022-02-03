package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func (c *gin.Context )  {
		c.JSON(http.StatusOK, gin.H{
			"status" : true,
			"messega" : "Welcome to web service Golang with gin",
		})
	})

	router.Run(":8000")
}