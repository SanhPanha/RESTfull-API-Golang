package services

import (
	models "book-author-api/domain"
	"book-author-api/pkg/repositories"
	"errors"
)


type AuthorService struct {
	repo repositories.IAuthorRepository
}

func NewAuthorService(repo repositories.IAuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) Create(author *models.Author) error {
    if s.repo == nil {
        return errors.New("repository not initialized")
    }
    return s.repo.Create(author)
}

func (s *AuthorService) GetAll() ([]models.Author, error) {
    if s.repo == nil {
        return nil, errors.New("repository not initialized")
    }
    return s.repo.GetAll()
}

func (s *AuthorService) GetByID(id uint) (*models.Author, error) {
    if s.repo == nil {
        return nil, errors.New("repository not initialized")
    }

    author, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    // Fetch book titles for additional processing or logging
    bookTitles, err := s.repo.GetBookTitlesByAuthorID(id)
    if err != nil {
        return nil, err
    }

    // You can use bookTitles here as needed
    _ = bookTitles // Example: Just using the variable to avoid compiler warning

    return author, nil
}


// func (s *AuthorService) GetByID(id uint) (*models.Author, error) {
//     if s.repo == nil {
//         return nil, errors.New("repository not initialized")
//     }

//     author, err := s.repo.GetByID(id)
//     if err != nil {
//         return nil, err
//     }

//     // Populate book titles
//     bookTitles, err := s.repo.GetBookTitlesByAuthorID(id)
//     if err != nil {
//         return nil, err
//     }

//     // Assign book titles to the author model
//     author.BookTitles = bookTitles

//     return author, nil
// }

func (s *AuthorService) Update(author *models.Author) error {
    if s.repo == nil {
        return errors.New("repository not initialized")
    }
    return s.repo.Update(author)
}

func (s *AuthorService) Delete(id uint) error {
    if s.repo == nil {
        return errors.New("repository not initialized")
    }
    return s.repo.Delete(id)
}