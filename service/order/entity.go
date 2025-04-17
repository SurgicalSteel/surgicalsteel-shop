package order

import (
	"surgicalsteel-shop/repository/order_item"
	"time"
)

type Order struct {
	Id         int64                  `json:"id"`
	CustomerId int64                  `json:"customer_id"`
	CreatedAt  time.Time              `json:"created_at"`
	Items      []order_item.OrderItem `json:"items"`
}
