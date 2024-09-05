package models

import "github.com/jinzhu/gorm"

type Book struct {
    gorm.Model
    Title    string
    AuthorID uint
    Author   Author
}
