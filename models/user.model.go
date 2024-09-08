package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255" json:"name"`
}
