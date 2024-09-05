package main

import (
    "github.com/gin-gonic/gin"
    "book-author-api/db"
    "book-author-api/routes"
)

func main() {
    db.InitDB()
    defer db.CloseDB()

    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
