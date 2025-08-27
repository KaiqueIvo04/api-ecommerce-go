package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.New(user.GetName(), user.GetEmail(), user.GetPassword(), user.GetType())
	log.Printf("User created: %+v\n", user)
	c.JSON(http.StatusCreated, gin.H{"user": user.GetName()})
}