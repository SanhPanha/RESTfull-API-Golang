package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "book-author-api/models"
    "book-author-api/services"
)

type AuthorController struct {
    AuthorService services.IAuthorService
}

// Controller methods

func (ctrl *AuthorController) CreateAuthor(c *gin.Context) {
    var author models.Author
    if err := c.ShouldBindJSON(&author); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.AuthorService.Create(&author); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, author)
}

func (ctrl *AuthorController) GetAuthors(c *gin.Context) {
    authors, err := ctrl.AuthorService.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, authors)
}

func (ctrl *AuthorController) GetAuthor(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
        return
    }
    author, err := ctrl.AuthorService.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
        return
    }
    c.JSON(http.StatusOK, author)
}

func (ctrl *AuthorController) UpdateAuthor(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
        return
    }
    var author models.Author
    if err := c.ShouldBindJSON(&author); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    author.ID = uint(id)
    if err := ctrl.AuthorService.Update(&author); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, author)
}

func (ctrl *AuthorController) DeleteAuthor(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
        return
    }
    if err := ctrl.AuthorService.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
