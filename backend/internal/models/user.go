package models

import (
	"errors"

	"gorm.io/gorm"
)
type User struct {
	gorm.Model
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
    PhoneNumber string `json:"phonenumber"`
	
}

func AddUser(db *gorm.DB, user User)error{
	result:=db.Create((&user))
	return result.Error
}
func GetUserPassword(db *gorm.DB, email string)(string, error){
	var user User
	if err:=db.Where("email=?", email).First(&user).Error; err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return "", errors.New("invalid email/password")
		}
		return "", err
	}
	
	return user.Password, nil
}