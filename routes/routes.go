package routes

import (
    "github.com/gin-gonic/gin"
    "book-author-api/controllers"
    "book-author-api/services"
)

func SetupRoutes(r *gin.Engine) {
    // Create services
    bookService := &services.BookService{}
    authorService := &services.AuthorService{}

    // Create controllers
    bookController := &controllers.BookController{BookService: bookService}
    authorController := &controllers.AuthorController{AuthorService: authorService}

    // Define routes
    r.GET("/books", bookController.GetBooks)
    r.POST("/books", bookController.CreateBook)
    r.GET("/books/:id", bookController.GetBook)
    r.PUT("/books/:id", bookController.UpdateBook)
    r.DELETE("/books/:id", bookController.DeleteBook)

    r.GET("/authors", authorController.GetAuthors)
    r.POST("/authors", authorController.CreateAuthor)
    r.GET("/authors/:id", authorController.GetAuthor)
    r.PUT("/authors/:id", authorController.UpdateAuthor)
    r.DELETE("/authors/:id", authorController.DeleteAuthor)
}
