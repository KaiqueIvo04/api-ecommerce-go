package main

import (
	"log"
	"os"

	"github.com/KaiqueIvo04/api-ecommerce-go/db"
	"github.com/KaiqueIvo04/api-ecommerce-go/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	dbConnectionURI := os.Getenv("DB_CONNECTION_URI")
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("environment variables could not be loaded: %s", err)
	}

	g := gin.Default()
	dbConnection, err := db.NewConnection(dbConnectionURI)
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	} else {
		log.Fatalf("Database connected!")
	}
	g.
	g.POST("/users", controller.CreateUser)

	g.Run(":8080")
}
