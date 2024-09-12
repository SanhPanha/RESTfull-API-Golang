package repositories

import (
	"book-author-api/db"
	models "book-author-api/domain"

	"gorm.io/gorm"
)

// Define the interface for AuthorRepository
type IAuthorRepository interface {
	Create(author *models.Author) error
	GetAll() ([]models.Author, error)
	GetByID(id uint) (*models.Author, error)
	Update(author *models.Author) error
	Delete(id uint) error
	GetBookTitlesByAuthorID(authorID uint) ([]string, error)
}

type AuthorRepository struct{}

func (r *AuthorRepository) Create(author *models.Author) error {
	return db.DB.Create(author).Error
}

func (r *AuthorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author
	err := db.DB.Preload("Books", func(db *gorm.DB) *gorm.DB {
		return db.Select("title")
	}).Find(&authors).Error
	return authors, err
}

// func (r *AuthorRepository) GetByID(id uint) (*models.Author, error) {
// 	var author models.Author
// 	err := db.DB.Preload("Books", func(db *gorm.DB) *gorm.DB {
// 		return db.Select("title")
// 	}).First(&author, id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &author, nil
// }
func (r *AuthorRepository) GetByID(id uint) (*models.Author, error) {
	var author models.Author
	err := db.DB.Preload("Books").First(&author, id).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}


func (r *AuthorRepository) Update(author *models.Author) error {
	return db.DB.Save(author).Error
}

func (r *AuthorRepository) Delete(id uint) error {
	return db.DB.Delete(&models.Author{}, id).Error
}

func (r *AuthorRepository) GetBookTitlesByAuthorID(authorID uint) ([]string, error) {
	var bookTitles []string
	err := db.DB.Model(&models.Book{}).
		Where("author_id = ?", authorID).
		Pluck("title", &bookTitles).Error
	return bookTitles, err
}