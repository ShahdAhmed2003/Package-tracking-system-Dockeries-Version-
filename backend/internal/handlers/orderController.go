package handlers

import (
	"bosta-backend/internal/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//GET /api/orders/addOrder
func CreateOrder(db *gorm.DB)http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var order models.Order
		if err:=json.NewDecoder(req.Body).Decode(&order); err!=nil{
			http.Error(res, "invalid request, please try again later", http.StatusBadRequest)
			return
		}
		order.Status="Pending"
		var user models.User
		if err:=db.Where("id=? and is_logged_in=true", order.UserId).First(&user).Error; err!=nil{
			if err==gorm.ErrRecordNotFound{
				http.Error(res, "You must login!", http.StatusUnauthorized)
			}else{
				http.Error(res, "error in retrieving from db!", http.StatusNotFound)
			}
			return
		}
		
		if order.UserId == "" || order.PickUpLocation.City == "" || order.DropOffLocation.City == "" {
			http.Error(res, "Missing required order fields", http.StatusBadRequest)
			return
		}

		if err:=db.Create(&order).Error; err!=nil{
			http.Error(res, "failed to create the order", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(order)
	}
}

//GET /api/orders/verify?orderId=123
func VerifyOrder(db *gorm.DB)http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		orderID:=req.URL.Query().Get("orderId")
		if orderID==""{
			http.Error(res, "order id is required", http.StatusBadRequest)
			return
		}
		var order models.Order
		if err:=db.First(&order, orderID).Error; err!=nil{
			if err==gorm.ErrRecordNotFound{
				http.Error(res, "order is not found", http.StatusNotFound)
			}else{
				http.Error(res, "failure in db", http.StatusNotFound)
			}
			return
		}

		order.Status = "Verified"
		if err := db.Save(&order).Error; err != nil {
			http.Error(res, "failed to update order status", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(order)
	}
}

//http://localhost:8080/api/users/orders?userId=1
func GetUserOrders(db *gorm.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		userID := req.URL.Query().Get("userId")
		if userID == "" {
			http.Error(res, "userId is required", http.StatusBadRequest)
			return
		}
		var user models.User
		if err:=db.Where("id=? and is_logged_in=true", userID).First(&user).Error; err!=nil{
			if err==gorm.ErrRecordNotFound{
				http.Error(res, "You must login!", http.StatusUnauthorized)
			}else{
				http.Error(res, "error in retrieving from db!", http.StatusNotFound)
			}
			return
		}

		var orders []models.Order
		if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
			http.Error(res, "failed to retrieve orders", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(orders)
	}
}


//http://localhost:8080/api/orders/details/1?userId=1
func GetOrderDetails(db *gorm.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		vars:=mux.Vars(req)
		orderID := vars["orderID"]
		userID := req.URL.Query().Get("userId")
		if orderID == "" {
			http.Error(res, "OrderID is required", http.StatusBadRequest)
			return
		}
		if userID == "" {
			http.Error(res, "userID is required", http.StatusBadRequest)
			return
		}

		var order models.Order
		if err := db.Where("id=? and user_id=?", orderID, userID).First(&order).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(res, "Order is not found", http.StatusNotFound)
			}else{
				http.Error(res, "Failed to retrieve order details", http.StatusInternalServerError)
			}
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(order)
	}
}


//
func CancelOrder(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//to get the order ID from the url
		vars:=mux.Vars(req)
		orderID := vars["orderID"]
		var order models.Order
		result := db.First(&order, orderID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Order not found", http.StatusNotFound)
			}else{
				http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			}
			return
		}
		if order.Status != "Pending" {
			http.Error(w, "Order cannot be canceled, it is not pending", http.StatusConflict)
			return
		}
		order.Status = "canceled"
		if err := db.Save(&order).Error; err != nil {
			http.Error(w, "Failed to cancel the order, try again later", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	}
}