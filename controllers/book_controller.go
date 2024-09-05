package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "book-author-api/db"
    "book-author-api/models"
)

func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&book)
    c.JSON(http.StatusCreated, book)
}

func GetBooks(c *gin.Context) {
    var books []models.Book
    db.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
    var book models.Book
    id := c.Param("id")
    if err := db.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
    var book models.Book
    id := c.Param("id")
    if err := db.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&book)
    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    if err := db.DB.Delete(&models.Book{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}