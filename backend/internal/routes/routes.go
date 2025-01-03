package routes

import (
	"bosta-backend/internal/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
    r := mux.NewRouter()
    r.Use(middlewares.CORSMiddleware) // Apply middleware here to all routes

    // User and Order routes
    UserRoutes(r, db)
    OrderRoutes(r, db)
    return r
}
