package repositories

import (
	"book-author-api/db"
	models "book-author-api/domain"
)

// Define the interface for BookRepository
type IBookRepository interface {
	Create(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
}

type BookRepository struct{}

func (r *BookRepository) Create(book *models.Book) error {
    return db.DB.Create(book).Error
}


func (r *BookRepository) GetAll() ([]models.Book, error) {
    var books []models.Book
    err := db.DB.Preload("Author").Find(&books).Error
    return books, err
}

func (r *BookRepository) GetByID(id uint) (*models.Book, error) {
    var book models.Book
    err := db.DB.Preload("Author").First(&book, id).Error
    if err != nil {
        return nil, err
    }
    return &book, nil
}



func (r *BookRepository) Update(book *models.Book) error {
	return db.DB.Save(book).Error
}

func (r *BookRepository) Delete(id uint) error {
	return db.DB.Delete(&models.Book{}, id).Error
}
