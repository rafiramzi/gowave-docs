package main

import (
	"log"
	"os"

	"gowave/database"
	"gowave/middlewares"
	"gowave/routes"
	"gowave/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		generate()
		return 
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("DB_NAME:", os.Getenv("DB_NAME"))

	// Initialize database connection
	db, err := database.Init()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Initialize Echo
	e := echo.New()

	e.Use(middlewares.ResponseMiddleware)

	// Register routes
	routes.RegisterRoutes(e, db)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func generate() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: go run main.go generate <folderName> <fileName>")
	}

	folderName := os.Args[2]
	fileName := os.Args[3]

	utils.Generate(folderName, fileName)
}
