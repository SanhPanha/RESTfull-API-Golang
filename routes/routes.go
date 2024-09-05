package routes

import (
    "github.com/gin-gonic/gin"
    "book-author-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/books", controllers.GetBooks)
    r.POST("/books", controllers.CreateBook)
    r.GET("/books/:id", controllers.GetBook)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)

    r.GET("/authors", controllers.GetAuthors)
    r.POST("/authors", controllers.CreateAuthor)
    r.GET("/authors/:id", controllers.GetAuthor)
    r.PUT("/authors/:id", controllers.UpdateAuthor)
    r.DELETE("/authors/:id", controllers.DeleteAuthor)
}
