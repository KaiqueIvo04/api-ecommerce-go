package userController

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
	"github.com/KaiqueIvo04/api-ecommerce-go/pkg/port"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService port.IUserService
}

func New(userService port.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	// Verify match of body with model
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Defines timeout for service operation
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	// Executes service operation
	response, err := uc.userService.Create(ctxTimeout, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Print(user)
	// Returns data
	ctx.JSON(http.StatusCreated, response)
}
