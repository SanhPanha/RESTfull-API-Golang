package models

import "github.com/jinzhu/gorm"

type Author struct {
    gorm.Model
    Name  string
    Books []Book
}
