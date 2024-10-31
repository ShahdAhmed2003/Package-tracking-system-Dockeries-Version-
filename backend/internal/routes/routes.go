package routes

import (
	"database/sql"
	"net/http"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) http.Handler {
    r := mux.NewRouter()

    UserRoutes(r, db)

    headers := gorillahandlers.AllowedHeaders([]string{"Content-Type"})
    methods := gorillahandlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"})
    origins := gorillahandlers.AllowedOrigins([]string{"http://localhost:3000"})

    return gorillahandlers.CORS(origins, headers, methods)(r)
}
