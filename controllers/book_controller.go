package controllers

import (
	"net/http"
	"strconv"

	models "book-author-api/domain"
	"book-author-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *services.BookService
}

// CreateBook handles the creation of a new book
func (ctrl *BookController) CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Ensure AuthorID is set correctly
    if book.AuthorID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "AuthorID is required"})
        return
    }

    if err := ctrl.BookService.Create(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, book)
}

// GetBooks retrieves all books
func (ctrl *BookController) GetBooks(c *gin.Context) {
	books, err := ctrl.BookService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBook retrieves a single book by ID
func (ctrl *BookController) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := ctrl.BookService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates a book's information by ID
func (ctrl *BookController) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Set the ID in the book to match the ID in the URL
	book.ID = uint(id)

	// Attempt to get the book to see if it exists before updating
	if _, err := ctrl.BookService.GetByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := ctrl.BookService.Update(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by ID
func (ctrl *BookController) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if _, err := ctrl.BookService.GetByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := ctrl.BookService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
