package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	// PostgreSQL connection string format:
	// "host=localhost port=5432 user=postgres dbname=bookstore password=yourpassword sslmode=disable"
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bookstore password=1234 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
