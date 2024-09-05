package db

import (
    "github.com/jinzhu/gorm"
    "book-author-api/models"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect to database")
    }
    DB.AutoMigrate(&models.Book{}, &models.Author{})
}

func CloseDB() {
    DB.Close()
}
