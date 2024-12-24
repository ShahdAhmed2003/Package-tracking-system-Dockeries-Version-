// main.go
package main

import (
	"bosta-backend/internal/database"
	"bosta-backend/internal/middlewares"
	"bosta-backend/internal/models"
	"bosta-backend/internal/routes"
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

    mux := routes.SetupRouter(db)
    handler := middlewares.CORSMiddleware(mux)
    log.Fatal(http.ListenAndServe(":8080", handler))
}
