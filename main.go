package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/customer/:id", customerHandler)

	router.GET("/query", queryHandler)
	router.Run(":8000")
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status" : true,
		"messega" : "Welcome to web service Golang with gin",
	})
}

func customerHandler(c *gin.Context){
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"customer_id" : id,
		"customer_name" : "Asrul Cahyadi Putra",
	})
}

func queryHandler(c *gin.Context){
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"customer_id" : "CUS-01",
		"customer_name" : name,
	})
}