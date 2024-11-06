package routes

import (
	"bosta-backend/internal/handlers"
	"bosta-backend/internal/middlewares"
	"net/http"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func OrderRoutes(req *mux.Router, db *gorm.DB){
	req.HandleFunc("/api/orders/addOrder", handlers.CreateOrder(db)).Methods("POST")
	req.HandleFunc("/api/orders/verify", handlers.VerifyOrder(db)).Methods("POST")
	req.HandleFunc("/api/users/orders", handlers.GetUserOrders(db)).Methods("GET")
	req.HandleFunc("/api/orders/details/{orderID}", handlers.GetOrderDetails(db)).Methods("GET")
	req.HandleFunc("/api/orders/cancel/{orderID}", handlers.CancelOrder(db)).Methods("DELETE")
	req.Handle("/api/orders/assign/{orderId}/{courierId}", middlewares.AuthMiddleware(http.HandlerFunc(handlers.AssignOrderToCourierHandler(db)))).Methods("POST")
	req.Handle("/api/orders/assigned_orders", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetAssignedOrders(db)))).Methods("GET")
	req.Handle("/api/orders/update-order-status/{order_id}", middlewares.AuthMiddleware(http.HandlerFunc(handlers.UpdateOrderStatus(db)))).Methods("PUT")
	req.Handle("/api/orders/courier/{courierId}",middlewares.AuthMiddleware( http.HandlerFunc(handlers.GetCourierOrders(db)))).Methods("GET")
	req.Handle("/api/orders/admin",(handlers.GetAllOrdersForAdmin(db))).Methods("GET")
	req.Handle("/api/orders/admin-update/{orderId}",(handlers.UpdateOrder(db))).Methods("PUT")
	req.Handle("/api/orders/admin-delete/{orderId}",(handlers.DeleteOrder(db))).Methods("DELETE")


	


}