package handlers

import (
	"fmt"
	"regexp"
)

func ValidateEmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}

func ValidatePhoneNumber(phoneNumber string) bool {
    re := regexp.MustCompile(`^[0-9]{10,15}$`)
    return re.MatchString(phoneNumber)
}

func ValidateUserData(email, name, password, phoneNumber string) error {
    if email == "" || !ValidateEmail(email) {
        return fmt.Errorf("a valid email is required")
    }
    if name == "" {
        return fmt.Errorf("name is required")
    }
    if password == "" {
        return fmt.Errorf("password is required")
    }
    if phoneNumber == "" || !ValidatePhoneNumber(phoneNumber) {
        return fmt.Errorf("a valid phone number is required (10-15 digits)")
    }
    return nil
}
