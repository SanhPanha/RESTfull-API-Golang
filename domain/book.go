package models

import "gorm.io/gorm"

// Book model
type Book struct {
    gorm.Model
    Title      string `json:"title"`
    AuthorID   uint   `json:"author_id"` // Ensure this is uint, not int
    AuthorName string `json:"author_name,omitempty" gorm:"-"`
    Author     Author `json:"author" gorm:"foreignKey:AuthorID"`
}

