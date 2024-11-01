package main

import (
	"bosta-backend/internal/database"
	"bosta-backend/internal/models"
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
    defer func(){
		sqlDB, _:=db.DB() //to get *sql.DB
		sqlDB.Close()
	}()

	//for auto migration
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("failed to migrate some tables: %v", err)
	}

	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("failed to migrate Order: %v", err)
	}

	err = db.AutoMigrate(&models.Address{})
	if err != nil {
		log.Fatalf("failed to migrate Address: %v", err)
	}

	err = db.AutoMigrate(&models.Package{})
	if err != nil {
		log.Fatalf("failed to migrate Package: %v", err)
	}


	  //initializing router with routes
    r := routes.SetupRouter(db)
    log.Fatal(http.ListenAndServe(":8080", r))
}
