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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbName == "" {
		log.Println("⚠️ Skipping database connection: missing DB_USER, DB_PASSWORD, or DB_NAME")
	} else {
		log.Println("DB_NAME:", dbName)

		db, err := database.Init()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		e := echo.New()
		e.Use(middlewares.ResponseMiddleware)
		routes.RegisterRoutes(e, db)
		port := os.Getenv("APP_PORT")
		if port == "" {
			port = "8080"
		}
		e.Logger.Fatal(e.Start(":" + port))
		return
	}
	e := echo.New()
	e.Use(middlewares.ResponseMiddleware)

	routes.RegisterRoutes(e, nil) 

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}


func generate() {
	if len(os.Args) < 4 {
		log.Fatalf("GoWave : Invalid arguments.")
	}

	folderName := os.Args[2]
	fileName := os.Args[3]

	utils.Generate(folderName, fileName)
}