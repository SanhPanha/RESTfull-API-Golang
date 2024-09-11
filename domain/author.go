package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"foreignKey:AuthorID"` // One-to-many relationship
}
