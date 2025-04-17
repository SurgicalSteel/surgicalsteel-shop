package order_item

type OrderItem struct {
	Id        int64 `json:"id"`
	OrderId   int64 `json:"order_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
