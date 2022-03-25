package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemsCode   string `json:"items_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
