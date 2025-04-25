package seeders

import (
	// "log"
	// "golang_projects/models"

	"gorm.io/gorm"
)

// Truncate all data from tables before seeding
func truncateTables(db *gorm.DB) {
	db.Exec("SET session_replication_role = 'replica';")
	
	//Example of how to truncate a table :

		// db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")   
}

// Seed data for Category, Supplier, and Product tables
func Seed(db *gorm.DB) {
	truncateTables(db)
	// Exmaple of how to seed a table :

		// seedUsers(db)
}

// Example create Seeder :

	// func seedUsers(db *gorm.DB) {
	// 	suppliers := []models.User{
	// 		{Username: "user1", Password: "password1"},
	// 	}

	// 	for _, supplier := range suppliers {
	// 		if err := db.Create(&supplier).Error; err != nil {
	// 			log.Printf("Could not seed suppliers: %v", err)
	// 		}
	// 	}
	// 	log.Println("Suppliers seeded")
	// }


