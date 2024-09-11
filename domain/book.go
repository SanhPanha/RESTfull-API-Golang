package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string
	AuthorID uint   // Foreign key to reference Author
	Author   Author `gorm:"foreignKey:AuthorID"` // Auto-loaded field for related Author
}
