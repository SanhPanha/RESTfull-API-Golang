package models

import "github.com/jinzhu/gorm"

// Author model with just a list of book titles, not the full Book structs
type Author struct {
	gorm.Model
	Name   string
	BookID uint
	Title  []string `gorm:"-"` // Virtual field (not stored in DB) for book titles, fetched in queries
}
