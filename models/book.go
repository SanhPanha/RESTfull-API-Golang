package models

import "github.com/jinzhu/gorm"

// Book model with only AuthorID and no direct reference to the Author struct
type Book struct {
	gorm.Model
	Title      string
	AuthorID   uint   // Foreign key to reference Author
	AuthorName string `gorm:"-"` // Virtual field (not stored in DB) for author name, fetched in queries
}
