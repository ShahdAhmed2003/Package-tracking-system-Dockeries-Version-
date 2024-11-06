package handlers

import (
	"bosta-backend/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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
func GetAssignedOrders(db*gorm.DB) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request){
	userClaims := r.Context().Value("user").(*models.Claims)
	var courier models.User
	if err := db.Where("email = ?", userClaims.Email).First(&courier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Courier not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}
	var orders []models.Order
	if err := db.Where("assigned_courier_id = ?", courier.ID).Find(&orders).Error; err != nil {
		http.Error(w, "Failed to retrieve assigned orders", http.StatusInternalServerError)
		return
	}
	if len(orders) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "No assigned orders found"})
		return
	}

	// Return the list of orders
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
}

func UpdateOrderStatus(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userClaims := r.Context().Value("user").(*models.Claims)

        // Retrieve the courier's ID using their email from the claims
        var courier models.User
        if err := db.Where("email = ?", userClaims.Email).First(&courier).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                http.Error(w, "Courier not found", http.StatusNotFound)
                return
            }
            http.Error(w, "Database error", http.StatusInternalServerError)
            return
        }

        // Get the order ID from the query parameters
        orderID := mux.Vars(r)["order_id"]

        if orderID == "" {
            http.Error(w, "Order ID is required", http.StatusBadRequest)
            return
        }

        // Retrieve the order to check its current assigned courier ID
        var order models.Order
        if err := db.Where("id = ?", orderID).First(&order).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                http.Error(w, "Order not found", http.StatusNotFound)
            } else {
                http.Error(w, "Database error", http.StatusInternalServerError)
            }
            return
        }

        // Check if the current courier is authorized to update this order
        if order.AssignedCourierID != courier.ID {
            http.Error(w, "Unauthorized to update this order", http.StatusForbidden)
            return
        }

        // Update the order status (assuming you get the new status from the request body)
        var updateRequest struct {
            Status string `json:"status"`
        }
        if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        // Update the order status
        order.Status = updateRequest.Status
        if err := db.Save(&order).Error; err != nil {
            http.Error(w, "Failed to update order status", http.StatusInternalServerError)
            return
        }

        // Return a success message
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Order status updated successfully"})
    }
}

func AssignOrderToCourierHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract the user claims from context to check the role
        userClaims := r.Context().Value("user").(*models.Claims)
        if userClaims.Role != "Admin" {
            http.Error(w, "Unauthorized: Admins only", http.StatusForbidden)
            return
        }

        // Parse the orderId and courierID from the request
        vars := mux.Vars(r)
        orderId, err := strconv.Atoi(vars["orderId"])
        if err != nil {
            http.Error(w, "Invalid order ID", http.StatusBadRequest)
            return
        }
        courierId, err := strconv.Atoi(vars["courierId"])
        if err != nil {
            http.Error(w, "Invalid courier ID", http.StatusBadRequest)
            return
        }

        // Check if the order exists
        var order models.Order
        if err := db.First(&order, orderId).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("order not found")
                http.Error(w, "Order not found", http.StatusNotFound)
            } else {
				log.Printf("DatabaseError")
                http.Error(w, "Database error", http.StatusInternalServerError)
            }
            return
        }

        // Check if the courier exists and has the "Courier" role
        var courier models.User
        if err := db.First(&courier, courierId).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("Courier not found")
                http.Error(w, "Courier not found", http.StatusNotFound)
            } else {
				log.Printf("DB error")
                http.Error(w, "Database error", http.StatusInternalServerError)
            }
            return
        }
        if courier.Role != "Courier" {
			log.Printf("The user is not a valid courie")
            http.Error(w, "The user is not a valid courier", http.StatusBadRequest)
            return
        }

        // Assign the order to the courier
        order.AssignedCourierID = uint(courierId)
        if err := db.Save(&order).Error; err != nil {
			
			log.Printf("failed to assign order")
            http.Error(w, "Failed to assign order", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Order assigned successfully"})
    }
}
// GET /api/orders/courier/{courierId}
func GetCourierOrders(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract the user claims from context to check the role
        userClaims := r.Context().Value("user").(*models.Claims)
        if userClaims.Role != "Admin" {
            http.Error(w, "Unauthorized: Admins only", http.StatusForbidden)
            return
        }

        // Parse the courierId from the URL
        vars := mux.Vars(r)
        courierId, err := strconv.Atoi(vars["courierId"])
        if err != nil {
            http.Error(w, "Invalid courier ID", http.StatusBadRequest)
            return
        }
		
        // Check if the courier exists
        var courier models.User
        if err := db.First(&courier, courierId).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                http.Error(w, "Courier not found", http.StatusNotFound)
            } else {
                http.Error(w, "Database error", http.StatusInternalServerError)
            }
            return
        }
        if courier.Role != "Courier" {
            http.Error(w, "The user is not a valid courier", http.StatusBadRequest)
            return
        }

        // Retrieve orders assigned to the specified courier
        var orders []models.Order
        if err := db.Where("assigned_courier_id = ?", courierId).Find(&orders).Error; err != nil {
            http.Error(w, "Failed to retrieve courier orders", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(orders)
    }
}

func GetAllOrdersForAdmin(db *gorm.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
    // Retrieve all orders from the database
    var orders []models.Order
    if err := db.Find(&orders).Error; err != nil {
        http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
        return
    }

    // Send the list of orders as the response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(orders)
}
}
func UpdateOrder(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID, err := strconv.Atoi(vars["orderId"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    var updatedOrder models.Order
    if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }

    var order models.Order
    if err := db.First(&order, orderID).Error; err != nil {
        http.Error(w, "Order not found", http.StatusNotFound)
        return
    }

    order.Status = updatedOrder.Status

    if err := db.Save(&order).Error; err != nil {
        http.Error(w, "Failed to update order", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(order)
}
}
func DeleteOrder (db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    orderID, err := strconv.Atoi(vars["orderId"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    if err := db.Delete(&models.Order{}, orderID).Error; err != nil {
        http.Error(w, "Failed to delete order", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
}
}