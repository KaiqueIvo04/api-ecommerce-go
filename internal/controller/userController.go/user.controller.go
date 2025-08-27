package usercontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
	"github.com/KaiqueIvo04/api-ecommerce-go/pkg/port"
)

type UserController struct {
	userService port.IUserService
}

func New(userService port.IUserService) *UserController{
	return &UserController{
		userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context, user *models.User) {
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userService.Add(c, user)
	if err != nil {
		panic(err)
	}

	log.Printf("User created: %+v\n", user)
	c.JSON(http.StatusCreated, gin.H{"user": user.GetName()})
}