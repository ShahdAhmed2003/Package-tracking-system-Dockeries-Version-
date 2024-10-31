package handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

//helper function to validate email structure
func ValidateEmail(email string)bool{
    re:=regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}

//helper function to validate email structure
func ValidatePhoneNumber(phonenumber string)bool{
    re:=regexp.MustCompile(`^[0-9]{10,15}$`)
    return re.MatchString(phonenumber)
}

func ValidateUserData(email, name, password, phoneNumber string) error{
	if email == ""||!ValidateEmail(email) {
			return fmt.Errorf("A valid email is required", http.StatusBadRequest)
			
		}
		if name == "" {
			return fmt.Errorf("Name is required", http.StatusBadRequest)
		}
		if password == "" {
			return fmt.Errorf("Password is required", http.StatusBadRequest)
		}
		if phoneNumber == ""||!ValidatePhoneNumber(phoneNumber) {
			return fmt.Errorf("A valid phone number is required(10-15) digits", http.StatusBadRequest)
		}
		return nil
}