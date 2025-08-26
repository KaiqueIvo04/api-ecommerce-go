package main

import (
	"api-ecommerce-go/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", controller.CreateUser)

	r.Run(":8080")
}
