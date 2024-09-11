package db

import (
	models "book-author-api/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("db/book_author.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Automigrate models
    err = DB.AutoMigrate(&models.Author{}, &models.Book{})
    if err != nil {
        panic("failed to automigrate models")
    }
}

// CloseDB closes the database connection
func CloseDB() {
    sqlDB, err := DB.DB()
    if err != nil {
        panic("failed to get SQL DB")
    }
    err = sqlDB.Close()
    if err != nil {
        panic("failed to close database")
    }
}
