package database

import (
	"fmt"

	"github.com/Schaitanya535/tailored_be/m/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=1234 dbname=tailored_be port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = connection
	fmt.Println("Database connected")
	connection.AutoMigrate(&models.User{}, &models.Product{})
	fmt.Println("Database migrated")

}
