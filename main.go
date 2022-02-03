package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/customer/:id", customerHandler)

	router.GET("/query", queryHandler)

	router.POST("/customer", postCustomerHandler)

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
		"status" : true,
		"message" : "record found!",
		"customer_id" : "CUS-01",
		"customer_name" : name,
	})
}

type InputCustomer struct {
	Name string 		`json:"name" binding:"required"`
	Age json.Number		`json:"age" binding:"required,number"`
	ParentName string 	`json:"parent_name" binding:"required"`
}

func postCustomerHandler(c *gin.Context){
	var inputCustomer InputCustomer

	err := c.ShouldBindJSON(&inputCustomer)

	if err != nil {
		errorMessages := []string{} //slice
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("error on field %s, rules: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) //append to slice errorMessages
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : errorMessages,
		})		
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : true,
		"message" : "Post request successfully!",
		"name" : inputCustomer.Name,
		"parent_name" : inputCustomer.ParentName,
		"age" : inputCustomer.Age,
	})
}