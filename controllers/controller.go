package controllers

import (
	"k6/models"
	"k6/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var profile models.RequestData
	if err := c.BindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	services.Insert(profile)
	c.JSON(http.StatusOK, profile)
}

// func EchoHandler(c *gin.Context) {
// 	var requestData RequestData

// 	//Bind JSON request body to the RequestData struct
// 	if err := c.ShouldBindJSON(&requestData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Process the JSON data
// 	// replyData := ResponseData{
// 	//     Reply: "Pong",
// 	// }

// 	// Respond with JSON
// 	c.JSON(http.StatusOK, "pong")
// }
