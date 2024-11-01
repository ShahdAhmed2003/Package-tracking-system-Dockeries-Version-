package handlers

import (
	"bosta-backend/internal/models"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)
func SignupHandler(db *gorm.DB)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //400
			return
		}

		if user.Role == "" {
			user.Role = "customer"
		}
		////
		if err:=ValidateUserData(user.Email, user.Name, user.Password, user.PhoneNumber); err!=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err:=models.AddUser(db, user); err!= nil {
			if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {//unique violation code
                http.Error(w, "Email already exists, no duplicate emails allowed", http.StatusConflict) //409
                return
            }
			http.Error(w, err.Error(), http.StatusInternalServerError) //500
			return
		}

		w.WriteHeader(http.StatusCreated) //201: Created
	}
    
}

func LoginHandler(db *gorm.DB)http.HandlerFunc {
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
		var user models.User
		res:=db.Where("email=?", email).First((&user))

		if res.Error != nil {
			if res.Error==gorm.ErrRecordNotFound{
				http.Error(w, "invalid email or password, please try again", http.StatusUnauthorized)
			} else{
				http.Error(w, res.Error.Error(), http.StatusInternalServerError)
			}
			return
		}

		if user.Password!=password{
			http.Error(w, "invalid password", http.StatusUnauthorized)
			return
		}
		user.IsLoggedIn=true
		if err := db.Save(&user).Error; err != nil {
			http.Error(w, "failed to update login status", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		// w.Write(([]byte("logged in successfully!")))
	}
}
