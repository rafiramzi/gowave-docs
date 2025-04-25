package main

import (
	"log"
	"os"
	"golang_projects/database"
	"golang_projects/middlewares"
	"golang_projects/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
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

	// AutoMigrate models :

		// if err := db.AutoMigrate(&models.User{}); err != nil {
		// 	log.Fatalf("Failed to migrate database: %v", err)
		// }

		// seeders.Seed(db)
		// seeders.SeedUsers(db) // Seeder for user

	// Initialize Echo
	e := echo.New()

	e.Use(middlewares.ResponseMiddleware)

	// Register routes
	routes.RegisterRoutes(e, db)    // Route untuk login

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}
	e.Logger.Fatal(e.Start(":" + port))
}
