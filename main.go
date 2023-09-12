package main

import (
	"k6/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/pingpong", controllers.Create)

	r.Run(":8080")
}
