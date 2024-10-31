package routes

import (
	"bosta-backend/internal/handlers"
	"database/sql"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router, db *sql.DB){
	r.HandleFunc("/signup", handlers.SignupHandler(db)).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
}