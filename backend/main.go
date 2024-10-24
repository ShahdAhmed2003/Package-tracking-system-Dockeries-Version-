package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers" // Package to handle CORS (added this)
	"github.com/gorilla/mux"      // Package for routing HTTP requests
	_ "github.com/lib/pq"         // PostgreSQL driver
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "asdqwe123"
    DB_NAME     = "myapp"
)

var db *sql.DB

type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    PhoneNumber string `json:"phonenumber"`
  
}

func init(){ //called when program starts
    var err error
	
    // Connect to the database using the connection string
    connStr := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err) // If there's an error, log it and stop the program
    }
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest) // 400: Bad Request
        return
    }

    // Insert the new user into the 'users' table
    _, err := db.Exec("INSERT INTO users(name, email, password,phonenumber) VALUES($1, $2, $3, $4)", user.Name, user.Email, user.Password,user.PhoneNumber)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) // 500: Server Error
        return
    }

    w.WriteHeader(http.StatusCreated) // 201: Created
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var req map[string]string
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    email:=req["email"]
    password:=req["password"]
    var userPass string
    err := db.QueryRow("select password from users where email=$1", email).Scan(&userPass)
    if err != nil {
        if err==sql.ErrNoRows{
            http.Error(w, "invalid email/password", http.StatusUnauthorized)
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


func main() {
    r := mux.NewRouter()  // Create a new router
    r.HandleFunc("/signup", signupHandler).Methods("POST") // Signup route (POST request)
    r.HandleFunc("/login", loginHandler).Methods("POST")

    // Handling CORS to allow requests from the React frontend
    headers := handlers.AllowedHeaders([]string{"Content-Type"}) // Allow 'Content-Type' header
    methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"}) // Allow POST, GET, and OPTIONS methods
    origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})  // Allow requests from this origin (React app running at localhost:3000)

    // Run the HTTP server on port 8080 with CORS enabled for the router 'r'
    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods)(r)))
}
