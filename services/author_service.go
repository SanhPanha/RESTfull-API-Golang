package services

import (
	"book-author-api/db"
	"book-author-api/models"
)

// Define the interface for AuthorService
type IAuthorService interface {
	Create(author *models.Author) error
	GetAll() ([]models.Author, error)
	GetByID(id uint) (*models.Author, error)
	Update(author *models.Author) error
	Delete(id uint) error
}

type AuthorService struct{}

func (s *AuthorService) Create(author *models.Author) error {
	return db.DB.Create(&author).Error
}

func (s *AuthorService) GetAll() ([]models.Author, error) {
	var authors []models.Author
	err := db.DB.Preload("Books").Find(&authors).Error
	return authors, err
}

func (s *AuthorService) GetByID(id uint) (*models.Author, error) {
	var author models.Author
	err := db.DB.First(&author, id).Error
	if err != nil {
		return nil, err
	}

	// Populate book titles in one go
	var bookTitles []string
	db.DB.Model(&models.Book{}).Where("author_id = ?", id).Pluck("title", &bookTitles)
	author.Title = bookTitles

	return &author, nil
}

func (s *AuthorService) Update(author *models.Author) error {
	return db.DB.Save(&author).Error
}

func (s *AuthorService) Delete(id uint) error {
	return db.DB.Delete(&models.Author{}, id).Error
}
