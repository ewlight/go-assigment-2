package resources

import "time"

type InputOrder struct {
	Id           uint             `json:"order_id"`
	CustomerName string           `json:"customer_name"`
	OrderedAt    time.Time        `json:"ordered_at"`
	Items        []InputOrderItem `json:"items"`
}

type UpdateOrder struct {
	Id           uint             `json:"order_id"  binding:"required"`
	CustomerName string           `json:"customer_name"`
	OrderedAt    time.Time        `json:"ordered_at"`
	Items        []InputOrderItem `json:"items"`
}

type InputOrderItem struct {
	ItemID      uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}