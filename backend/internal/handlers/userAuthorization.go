package handlers

import (
	"bosta-backend/internal/models"
	"fmt"
	"net/http"
	"strconv" // Import strconv for integer-to-string conversion
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"userID": strconv.FormatUint(uint64(user.ID), 10),
		"email":  user.Email,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("your_secret_key")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}


func VerifyToken(req *http.Request) (string, error) {
    authHeader := req.Header.Get("Authorization")
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        return "", fmt.Errorf("authorization token required")
    }
    
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte("your_secret_key"), nil
    })

    if err != nil || !token.Valid {
        return "", fmt.Errorf("invalid token")
    }
    
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || claims["userID"] == nil {
        return "", fmt.Errorf("user ID missing in token")
    }
    
    return claims["userID"].(string), nil
}
