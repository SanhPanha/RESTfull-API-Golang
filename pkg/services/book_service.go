package services

import (
	models "book-author-api/domain"
	"book-author-api/pkg/repositories"
)

// Define the interface for BookService

type BookService struct {
	repo repositories.IBookRepository
}

// NewBookService initializes a new BookService
func NewBookService(repo repositories.IBookRepository) *BookService {
	return &BookService{repo: repo}
}

// Create a book
func (s *BookService) Create(book *models.Book) error {
	
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