package models

import "github.com/jinzhu/gorm"

type Book struct {
    gorm.Model
    Title  string
    Author Author `gorm:"foreignkey:AuthorID"`
    AuthorID uint
}

type Author struct {
    gorm.Model
    Name  string
    Books []Book
}