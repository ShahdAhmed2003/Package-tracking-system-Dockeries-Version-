package main

import (
	"bosta-backend/internal/database"
	"bosta-backend/internal/routes"
	"log"
	"net/http"
)

func main() {
	  //initializing db connection
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }
    defer db.Close()

	  //initializing router with routes
    r := routes.SetupRouter(db)
    log.Fatal(http.ListenAndServe(":8080", r))
}
