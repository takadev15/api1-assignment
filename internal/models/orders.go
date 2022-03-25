package models

type Order struct {
	GormModel
	CustomerName string `json:"customer_name"`
	Item         Item   `gorm:"embedded"`
}
