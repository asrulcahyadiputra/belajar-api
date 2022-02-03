package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.Run(":8000")
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status" : true,
		"messega" : "Welcome to web service Golang with gin",
	})
}