package main

import (
	"context"
	"log"
	"os"

	"github.com/KaiqueIvo04/api-ecommerce-go/db"
	userController "github.com/KaiqueIvo04/api-ecommerce-go/internal/controllers"
	userRepository "github.com/KaiqueIvo04/api-ecommerce-go/internal/repositories/userRepository"
	userService "github.com/KaiqueIvo04/api-ecommerce-go/internal/services/userService"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	// Variáveis de ambiente
	err := godotenv.Load()
	dbConnectionURI := os.Getenv("DB_CONNECTION_URI")

	if err != nil {
		log.Fatalf("environment variables could not be loaded: %s", err)
	}

	// Setup Database
	dbConnection, err := db.NewConnection(ctx, dbConnectionURI)
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	} else {
		log.Println("Database connected!")
	}
	defer dbConnection.Close()

	// Setup server
	g := gin.Default()

	// User
	userRepository := userRepository.New(dbConnection)
	userService := userService.New(userRepository)
	userController := userController.New(userService)

	g.POST("/users", userController.CreateUser)

	g.Run(":8080")
}
