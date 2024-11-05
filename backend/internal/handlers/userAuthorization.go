package handlers

import (
	"bosta-backend/internal/models"
	"time"
	"github.com/dgrijalva/jwt-go"
	
)
func createToken(user models.User) (string, error) {
	claims := models.Claims{
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
		},		
	}

	// Create the token using your signing method (HS256 in this case)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("your_secret_key") // Use a secure method to store your secret key

	// Sign the token
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
