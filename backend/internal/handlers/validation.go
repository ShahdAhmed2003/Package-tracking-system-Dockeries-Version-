package handlers

import (
	"fmt"
	"regexp"
)

func ValidateEmail(email string)bool{
    re:=regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}

func ValidatePhoneNumber(phonenumber string)bool{
    re:=regexp.MustCompile(`^[0-9]{10,15}$`)
    return re.MatchString(phonenumber)
}

func ValidateUserData(email, name, password, phoneNumber string) error{
	if email == ""||!ValidateEmail(email) {
			return fmt.Errorf("A valid email is required")
			
		}
		if name == "" {
			return fmt.Errorf("Name is required")
		}
		if password == "" {
			return fmt.Errorf("Password is required")
		}
		if phoneNumber == ""||!ValidatePhoneNumber(phoneNumber) {
			return fmt.Errorf("A valid phone number is required(10-15) digits")
		}
		return nil
}