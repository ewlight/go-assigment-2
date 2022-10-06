package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Id           uint      `json:"order_id" gorm:"primary_key"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
	Items        []Item    `gorm:"foreignKey:OrderID"`
}

type Item struct {
	gorm.Model
	Id          uint  `json:"item_id" gorm:"primary_key"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID 	uint  `json:"order_id"`
	Order   	Order `gorm:"foreignKey:OrderID"`
}

