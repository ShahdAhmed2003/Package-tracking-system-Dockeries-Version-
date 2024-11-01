package models

import (
	"time"

	"gorm.io/gorm"
)
type Order struct {
    gorm.Model
    UserId     string `json:"user_id"`

    //embedded tag allows me to treat the struct as part of the parent struct and flatten the fields
    PickUpLocation    Address `json:"pickup_location"  gorm:"embedded;embeddedPrefix:pickup_"`
    DropOffLocation Address `json:"drop_off_location" gorm:"embedded;embeddedPrefix:dropoff_"`
    PackageDetails Package `json:"package_details" gorm:"embedded;embeddedPrefix:package_"`
    CourierInfo string `json:"courier_info"`
    Status string `json:"status"`
    DeliveryTime *time.Time `json:"delivery_time"`
	
}
