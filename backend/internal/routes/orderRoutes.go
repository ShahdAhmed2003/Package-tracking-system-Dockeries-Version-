package routes

import (
	"bosta-backend/internal/handlers"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func OrderRoutes(req *mux.Router, db *gorm.DB){
	req.HandleFunc("/api/orders/addOrder", handlers.CreateOrder(db)).Methods("POST")
	req.HandleFunc("/api/orders/verify", handlers.VerifyOrder(db)).Methods("POST")
	req.HandleFunc("/api/users/orders", handlers.GetUserOrders(db)).Methods("GET")
	req.HandleFunc("/api/orders/details/{orderID}", handlers.GetOrderDetails(db)).Methods("GET")
	req.HandleFunc("/api/orders/cancel/{orderID}", handlers.CancelOrder(db)).Methods("DELETE")
}