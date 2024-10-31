package routes

import (
	"bosta-backend/internal/handlers"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router, db *gorm.DB){
	r.HandleFunc("/signup", handlers.SignupHandler(db)).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
}