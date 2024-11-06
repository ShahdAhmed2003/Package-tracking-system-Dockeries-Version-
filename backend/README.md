# Bosta API

This is the backend API for the Bosta application, built with Go and GORM for user management, order processing, and other services. The API handles user authentication, order creation and verification, and package handling for a delivery service.



## Getting Started

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-repo/bosta-backend.git
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure the database** in `database/config.go` or as environment variables.

4. **Run the API server**:
   ```bash
   go run main.go
   ```

## Database Models

### User Model
```go
type User struct {
    Name        string
    Email       string `gorm:"unique"`
    Password    string
    PhoneNumber string
    Role        string
    IsLoggedIn  bool   `gorm:"default:false"`
    Orders      []Order
}
```

### Order Model
```go
type Order struct {
    UserId         string
    PickUpLocation Address `gorm:"embedded;embeddedPrefix:pickup_"`
    DropOffLocation Address `gorm:"embedded;embeddedPrefix:dropoff_"`
    PackageDetails Package `gorm:"embedded;embeddedPrefix:package_"`
    CourierInfo    string
    Status         string
    DeliveryTime   *time.Time
}
```

### Address Model
```go
type Address struct {
    StreetAddress string
    City          string
    State         string
    PostalCode    string
    Country       string
}
```

## API Endpoints

### User Endpoints

#### 1. **Sign Up User**

   - **URL**: `/signup`
   - **Method**: `POST`
   - **Body**:
     ```json
     {
       "name": "mennatallah",
       "email": "m@gmail.com",
       "password": "password123",
       "phonenumber": "1234567890"
     }
     ```
   - **Response**:
     - `201 Created`: User successfully created.
     - `409 Conflict`: Email already exists.
     - `400 Bad Request`: Validation error.

#### 2. **Login User**

   - **URL**: `/login`
   - **Method**: `POST`
   - **Body**:
     ```json
     {
       "email": "m@gmail.com",
       "password": "password123"
     }
     ```
   - **Response**:
     - `200 OK`: Returns user data.
     - `401 Unauthorized`: Invalid email or password.
     - `400 Bad Request`: Missing email or password.

### Order Endpoints

#### 1. **Create Order**

   - **URL**: `/api/orders/addOrder`
   - **Method**: `POST`
   - **Body**:
     ```json
     {
       "user_id": "1",
       "pickup_location": {
         "street_address": "123 Main St",
         "city": "Cairo",
         "state": "CA",
         "postal_code": "90001",
         "country": "Egypt"
       },
       "drop_off_location": {
         "street_address": "456 Elm St",
         "city": "Giza",
         "state": "GZ",
         "postal_code": "80001",
         "country": "Egypt"
       },
       "package_details": {
         "weight": 2.5,
         "length": 10.0,
         "width": 5.0,
         "height": 2.0,
         "contents": "Documents",
         "is_fragile": false,
         "special_requirements": "None"
       },
       "courier_info": "Express"
     }
     ```
   - **Response**:
     - `201 Created`: Order created successfully.
     - `400 Bad Request`: Missing required order fields.

#### 2. **Verify Order**

   - **URL**: `/api/orders/verify`
   - **Method**: `GET`
   - **Query Params**: `orderId`
   - **Response**:
     - `200 OK`: Order status updated to "Verified".
     - `404 Not Found`: Order not found.
     - `400 Bad Request`: Missing `orderId`.

#### 3. **Get User Orders**

   - **URL**: `/api/users/orders`
   - **Method**: `GET`
   - **Query Params**: `userId`
   - **Response**:
     - `200 OK`: Returns list of user orders.
     - `401 Unauthorized`: User not logged in.
     - `400 Bad Request`: Missing `userId`.

#### 4. **Get Order Details**

   - **URL**: `/api/orders/details/{orderID}`
   - **Method**: `GET`
   - **Query Params**: `userId`
   - **Response**:
     - `200 OK`: Returns order details.
     - `404 Not Found`: Order or user not found.
     - `400 Bad Request`: Missing `orderID` or `userId`.

#### 5. **Cancel Order**

   - **URL**: `/api/orders/cancel/{orderID}`
   - **Method**: `POST`
   - **Response**:
     - `200 OK`: Order canceled successfully and status changed to "canceled".
     - `409 Conflict`: Order cannot be canceled if it’s not pending.
     - `404 Not Found`: Order not found.

## Error Handling

The API provides standard HTTP error codes to indicate errors:
- `400 Bad Request`: Invalid input or missing required fields.
- `401 Unauthorized`: Authentication error.
- `404 Not Found`: Resource not found.
- `409 Conflict`: Conflict error (e.g., duplicate entry).
- `500 Internal Server Error`: Server error.

## Technologies Used

- **Golang** for server-side programming
- **GORM** for ORM and database interactions
- **Gorilla Mux** for HTTP routing
