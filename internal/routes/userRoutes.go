package routes

import (
	"bosta-backend/internal/handlers"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router, db *gorm.DB){
	r.HandleFunc("/signup", handlers.SignupHandler(db)).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
} 
/*func AssignOrderToCourierHandler(db *gorm.DB) http.HandlerFunc {
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
                http.Error(w, "Order not found", http.StatusNotFound)
            } else {
                http.Error(w, "Database error", http.StatusInternalServerError)
            }
            return
        }

        // Check if the courier exists and has the "Courier" role
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

        // Assign the order to the courier
        order.AssignedCourierID = uint(courierId)
        if err := db.Save(&order).Error; err != nil {
            http.Error(w, "Failed to assign order", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Order assigned successfully"})
    }
}*/