package middlewares
import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"bosta-backend/internal/models"
)
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims := &models.Claims{} // Struct to hold the decoded JWT claims
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil // Replace with a secure method of retrieving your secret key
		})

		if err != nil || !tkn.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set claims in context
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}