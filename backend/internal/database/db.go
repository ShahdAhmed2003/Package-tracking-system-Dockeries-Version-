package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "Menna@123#"
    DB_NAME     = "bosta"
)

func ConnectDB()(*sql.DB, error){
	connStr := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("error connecting to the database:", err) // If there's an error, log it and stop the program
    }
	return db, err
}