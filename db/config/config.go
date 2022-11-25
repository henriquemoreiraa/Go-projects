package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	dns := "host=localhost user=postgres password=123456 dbname=book port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Printf("Database conected...")
	return db
}