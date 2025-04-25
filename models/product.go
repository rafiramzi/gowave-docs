package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	CategoryID  uint    `json:"category_id"`
	SupplierID  uint    `json:"supplier_id"`
	Category    Category
	Supplier    Supplier
}
