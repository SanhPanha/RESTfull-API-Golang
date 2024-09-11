package routes

import (
	"book-author-api/controllers"
	"book-author-api/pkg/repositories"
	"book-author-api/pkg/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Initialize repositories
	authorRepo := &repositories.AuthorRepository{}
	bookRepo := &repositories.BookRepository{}

	// Create services with repositories
	authorService := services.NewAuthorService(authorRepo)
	bookService := services.NewBookService(bookRepo)

	// Create controllers
	authorController := &controllers.AuthorController{AuthorService: authorService}
	bookController := &controllers.BookController{BookService: bookService}

	// Group book-related routes
	books := r.Group("/books")
	{
		books.GET("/", bookController.GetBooks)
		books.POST("/", bookController.CreateBook)
		books.GET("/:id", bookController.GetBook)
		books.PUT("/:id", bookController.UpdateBook)
		books.DELETE("/:id", bookController.DeleteBook)
	}

	// Group author-related routes
	authors := r.Group("/authors")
	{
		authors.GET("/", authorController.GetAuthors)  // Fetch all authors
		authors.POST("/", authorController.CreateAuthor)
		authors.GET("/:id", authorController.GetAuthor) // Fetch a specific author by ID
		authors.PUT("/:id", authorController.UpdateAuthor)
		authors.DELETE("/:id", authorController.DeleteAuthor)
    
	}
}
