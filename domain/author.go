package models

import (
	"gorm.io/gorm"
)

// Author model
type Author struct {
    gorm.Model
    Name       string `json:"name"`
    Books      []Book `json:"books" gorm:"foreignKey:AuthorID"`
    BookTitles []string `json:"book_titles" gorm:"-"`
}
