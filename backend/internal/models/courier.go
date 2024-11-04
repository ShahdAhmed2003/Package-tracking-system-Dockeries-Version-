package models



type Courier struct {
    CourierID uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"unique"`
    Orders     []Order  `gorm:"foreignKey:CourierID"` // Associate Orders with Courier
}
