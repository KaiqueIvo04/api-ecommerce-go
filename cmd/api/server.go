package main

import (

	"log"
	"os"

	"github.com/KaiqueIvo04/api-ecommerce-go/db"
	usercontroller "github.com/KaiqueIvo04/api-ecommerce-go/internal/controller/userController.go"
	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
	"github.com/KaiqueIvo04/api-ecommerce-go/internal/repository/userRepository"
	"github.com/KaiqueIvo04/api-ecommerce-go/internal/service/userService"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Variáveis de ambiente
	err := godotenv.Load()
	dbConnectionURI := os.Getenv("DB_CONNECTION_URI")

	if err != nil {
		log.Fatalf("environment variables could not be loaded: %s", err)
	}

	// Setup Database
	dbConnection, err := db.NewConnection(dbConnectionURI)
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	} else {
		log.Fatalf("Database connected!")
	}
	defer dbConnection.Close()

	// Setup server
	g := gin.Default()

	// User
	userRepository := userrepository.New(dbConnection)
	userService := userservice.New(userRepository)
	userController := usercontroller.New(userService)

	g.POST("/users", func(ctx *gin.Context) {
		var user models.User
		userController.CreateUser(ctx, &user)
	})

	g.Run(":8080")
}
