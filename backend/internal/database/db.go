package database

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB() (*gorm.DB, error) {
    // Read database connection details from environment variables
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Construct the Data Source Name (DSN)
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
        dbHost, dbPort, dbUser, dbPassword, dbName)

    // Open a connection to the PostgreSQL database using sql package
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    // Create the database if it doesn't exist
    _, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
    if err != nil && err.Error() != fmt.Sprintf("pq: database \"%s\" already exists", dbName) {
        return nil, err
    }

    // Connect to the new database using GORM
    newDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
        dbHost, dbPort, dbUser, dbPassword, dbName)
    gormDB, err := gorm.Open(postgres.Open(newDsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return gormDB, nil
}
