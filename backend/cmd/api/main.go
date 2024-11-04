// main.go
package main

import (
	"bosta-backend/internal/database"
	"bosta-backend/internal/middlewares"
	"bosta-backend/internal/models"
	"bosta-backend/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }
    defer func() {
        sqlDB, _ := db.DB()
        sqlDB.Close()
    }()

    err = db.AutoMigrate(&models.User{}, &models.Order{}, &models.Address{}, &models.Package{})
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    r := routes.SetupRouter(db)
    // Ensure middleware is applied correctly
    r.Use(middlewares.CORSMiddleware)

    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
