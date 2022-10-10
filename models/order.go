package models

import (
	"time"
)

type Order struct {
	ID            string    `json:"order_id" gorm:"primaryKey"`
	Customer_name string    `json:"customer_name"`
	Ordered_at    time.Time `json:"ordered_at"`
	Items         []Item    `json:"items" gorm:"foreignKey:Order_id"`
}
