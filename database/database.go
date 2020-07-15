package database

import (
	"github.com/jinzhu/gorm"
	"github.com/kingsman20/fintech_app/helpers"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function
func InitDatabase() {
	database, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=kingslife dbname=fintech password=kingslife sslmode=disable")
	helpers.HandleErr(err)
	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
