package models

import "time"

type Order struct {
	BaseModel
	CustomerName string    `json:"customer_name" gorm:"column:customerName;not null"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"column:orderedAt;not null"`
	Items        []Item    `json:"items,omitempty" gorm:"foreignkey:OrderId"`
}

type Item struct {
	BaseModel
	ItemCode    string `json:"item_code" gorm:"column:item_code;not null"`
	Description string `json:"description" gorm:"column:description;not null"`
	Quantity    int    `json:"quantity" gorm:"column:quantity;not null"`
	OrderId     uint   `json:"order_id,omitempty" gorm:"column:orderId;not null"`
}

type CreateOrderDto struct {
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []struct {
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}

type UpdateOrderDto struct {
	CustomerName string `json:"customer_name"`
	Items        []struct {
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}
