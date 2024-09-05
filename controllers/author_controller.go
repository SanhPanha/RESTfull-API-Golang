package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "book-author-api/db"
    "book-author-api/models"
)

func CreateAuthor(c *gin.Context) {
    var author models.Author
    if err := c.ShouldBindJSON(&author); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&author)
    c.JSON(http.StatusCreated, author)
}

func GetAuthors(c *gin.Context) {
    var authors []models.Author
    db.DB.Find(&authors)
    c.JSON(http.StatusOK, authors)
}

func GetAuthor(c *gin.Context) {
    var author models.Author
    id := c.Param("id")
    if err := db.DB.First(&author, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
        return
    }
    c.JSON(http.StatusOK, author)
}

func UpdateAuthor(c *gin.Context) {
    var author models.Author
    id := c.Param("id")
    if err := db.DB.First(&author, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
        return
    }
    if err := c.ShouldBindJSON(&author); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&author)
    c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {
    id := c.Param("id")
    if err := db.DB.Delete(&models.Author{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
