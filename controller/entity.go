package controller

type CreateOrderRequest struct {
}

type CreateProductRequest struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int64  `json:"price"`
	Stock int64  `json:"stock"`
}
