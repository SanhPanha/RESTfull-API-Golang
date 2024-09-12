package services

import (
	models "book-author-api/domain"
	"book-author-api/pkg/repositories"
	"errors"

	"gorm.io/gorm"
)

// Define the interface for BookService

type BookService struct {
    repo        repositories.IBookRepository
    authorRepo  repositories.IAuthorRepository
}

// NewBookService initializes a new BookService
func NewBookService(repo repositories.IBookRepository, authorRepo repositories.IAuthorRepository) *BookService {
    return &BookService{repo: repo, authorRepo: authorRepo}
}

func (s *BookService) AuthorExists(authorID uint) (bool, error) {
    _, err := s.authorRepo.GetByID(authorID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, nil
        }
        return false, err
    }
    return true, nil
}


// Create a book
func (s *BookService) Create(book *models.Book) error {
    // Ensure the author exists
    exists, err := s.AuthorExists(book.AuthorID)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("author does not exist")
    }
    
    return s.repo.Create(book)
}


// Get all books
func (s *BookService) GetAll() ([]models.Book, error) {
	return s.repo.GetAll()
}


// Get a book by ID
func (s *BookService) GetByID(id uint) (*models.Book, error) {
	return s.repo.GetByID(id)
}

// Update a book
func (s *BookService) Update(book *models.Book) error {
	return s.repo.Update(book)
}

// Delete a book by ID
func (s *BookService) Delete(id uint) error {
	return s.repo.Delete(id)
}