package services

import (
	"book-author-api/db"
	"book-author-api/models"
)

// Define other methods such as GetAll, GetByID, Update, Delete
type IBookService interface {
	Create(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
}

type BookService struct{}

func (s *BookService) Create(book *models.Book) error {
	return db.DB.Create(&book).Error
}

func (s *BookService) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := db.DB.Model(&models.Book{}).Select("books.*, authors.name as author_name").
		Joins("left join authors on authors.id = books.author_id").Scan(&books).Error
	return books, err
}

func (s *BookService) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := db.DB.Model(&models.Book{}).Select("books.*, authors.name as author_name").
		Joins("left join authors on authors.id = books.author_id").
		Where("books.id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) Update(book *models.Book) error {
	return db.DB.Save(&book).Error
}

func (s *BookService) Delete(id uint) error {
	return db.DB.Delete(&models.Book{}, id).Error
}
