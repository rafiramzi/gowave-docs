package models

import "github.com/jinzhu/gorm"

type Supplier struct {
	gorm.Model
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Products []Product `gorm:"foreignkey:SupplierID"`
}
