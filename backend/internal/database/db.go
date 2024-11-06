package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "asdqwe123"
	DB_NAME     = "bosta"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
    defer db.Close()

    //now it creates the db if it doesnt exist:
    _, err=db.Exec(fmt.Sprintf("create database %s", DB_NAME))
    if err!=nil&& err.Error()!="pq: database \"bosta\" already exists"{
        return nil, err
    }

    //connect to the new db:
	newDsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    gormDB, err:=gorm.Open(postgres.Open(newDsn), &gorm.Config{})
    if err!=nil{
        return nil, err
    }
	return gormDB, nil
}
