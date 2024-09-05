package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "book-author-api/models"  // Import the models package
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }

    // Automigrate models from the models package
    DB.AutoMigrate(&models.Book{}, &models.Author{})
}

func CloseDB() {
    DB.Close()
}