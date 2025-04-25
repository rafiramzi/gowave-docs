package seeders

import (
	"log"
	"math/rand"
	"golang_projects/models"
	"strconv"

	"gorm.io/gorm"
)

// Truncate all data from tables before seeding
func truncateTables(db *gorm.DB) {
	// Disable foreign key checks (if required for truncation)
	db.Exec("SET session_replication_role = 'replica';")

	// Truncate tables in a specific order to avoid foreign key constraint errors
	db.Exec("TRUNCATE TABLE products RESTART IDENTITY CASCADE;")
	db.Exec("TRUNCATE TABLE categories RESTART IDENTITY CASCADE;")
	db.Exec("TRUNCATE TABLE suppliers RESTART IDENTITY CASCADE;")

	// Re-enable foreign key checks (if previously disabled)
	db.Exec("SET session_replication_role = 'origin';")
}

// Seed data for Category, Supplier, and Product tables
func Seed(db *gorm.DB) {
	// Truncate tables before seeding
	truncateTables(db)

	seedCategories(db)
	seedSuppliers(db)
	seedProducts(db)
}

// Seed 5 categories
func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Electronics"},
		{Name: "Furniture"},
		{Name: "Clothing"},
		{Name: "Books"},
		{Name: "Sports Equipment"},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			log.Printf("Could not seed categories: %v", err)
		}
	}
	log.Println("Categories seeded")
}

// Seed 5 suppliers
func seedSuppliers(db *gorm.DB) {
	suppliers := []models.Supplier{
		{Name: "Supplier A", Address: "123 Supplier St."},
		{Name: "Supplier B", Address: "456 Supplier Ave."},
		{Name: "Supplier C", Address: "789 Supplier Blvd."},
		{Name: "Supplier D", Address: "101 Supplier Ln."},
		{Name: "Supplier E", Address: "202 Supplier Dr."},
	}

	for _, supplier := range suppliers {
		if err := db.Create(&supplier).Error; err != nil {
			log.Printf("Could not seed suppliers: %v", err)
		}
	}
	log.Println("Suppliers seeded")
}

// Seed 20 products
func seedProducts(db *gorm.DB) {
	var categories []models.Category
	var suppliers []models.Supplier

	// Fetch categories and suppliers to assign to products
	db.Find(&categories)
	db.Find(&suppliers)

	products := []models.Product{}

	for i := 1; i <= 20; i++ {
		product := models.Product{
			Name:        "Product " + strconv.Itoa(i), // Using strconv.Itoa to convert int to string
			Price:       float64(rand.Intn(100) + 1),
			Description: "Description for product " + strconv.Itoa(i),
			Stock:       rand.Intn(50) + 1,
			CategoryID:  categories[rand.Intn(len(categories))].ID,
			SupplierID:  suppliers[rand.Intn(len(suppliers))].ID,
		}
		products = append(products, product)
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			log.Printf("Could not seed products: %v", err)
		}
	}
	log.Println("Products seeded")
}
