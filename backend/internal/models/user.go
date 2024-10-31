package models

import (
	"database/sql"
	"errors"
)
type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    PhoneNumber string `json:"phonenumber"`
  
}

func AddUser(db *sql.DB, user User)error{
    _, err := db.Exec("INSERT INTO users(name, email, password, phonenumber) VALUES($1, $2, $3, $4)", user.Name, user.Email, user.Password, user.PhoneNumber)
	return err
}
func GetUserPassword(db *sql.DB, email string)(string, error){
	var password string
	err:=db.QueryRow("select password from users where email=$1", email).Scan(&password)
	if err==sql.ErrNoRows{
		return "", errors.New("invalid email/password")
	}
	return password, err
}