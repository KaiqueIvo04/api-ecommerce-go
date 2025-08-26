package controller

import (
	"api-ecommerce-go/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.New(user.GetName(), user.GetEmail(), user.GetPassword(), user.GetType())
	fmt.Printf("User created: %+v\n", user)
	c.JSON(http.StatusCreated, gin.H{"user": user.GetName()})
}