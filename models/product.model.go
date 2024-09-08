package models

import "gorm.io/gorm"

type Product struct {
	Model       gorm.Model
	Name        string `gorm:"size:255" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Price       int    `json:"price"`
}
