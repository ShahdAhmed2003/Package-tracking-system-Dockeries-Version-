package handlers

import (
	"bosta-backend/internal/models"
	"database/sql"
	"encoding/json"
	"net/http"
)
func SignupHandler(db *sql.DB)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //400: Bad Request
			return
		}

		////
		if err:=ValidateUserData(user.Email, user.Name, user.Password, user.PhoneNumber); err!=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err:=models.AddUser(db, user); err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // 409: Conflict
			return
		}

		w.WriteHeader(http.StatusCreated) //201: Created
	}
    
}

func LoginHandler(db *sql.DB)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email:=req["email"]
		password:=req["password"]
		if email==""{
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}
		if password==""{
			http.Error(w, "Password is required", http.StatusBadRequest)
			return
		}
		var userPass string
		err := db.QueryRow("select password from users where email=$1", email).Scan(&userPass)
		if err != nil {
			if err==sql.ErrNoRows{
				http.Error(w, "invalid email or password, please try again", http.StatusUnauthorized)
			} else{
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if userPass!=password{
			http.Error(w, "invalid password", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(([]byte("logged in successfully!")))
	}
}
