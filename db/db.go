package db

import (
    "log"
	"github.com/26thavenue/creditCardValidator/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "fmt"

)


var DB *gorm.DB

func init() {
    // Initialize SQLite database
    var err error
    DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Migrate the database schema
    err = DB.AutoMigrate(&models.User{}, &models.CreditCard{})
    if err != nil {
        log.Fatalf("Failed to migrate the database schema: %v", err)
    }

    fmt.Println("Database connection successful")
}
