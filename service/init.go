package service

import (
	"surgicalsteel-shop/service/inventory"
	"surgicalsteel-shop/service/order"
	"surgicalsteel-shop/service/product"
	"surgicalsteel-shop/service/quote"
)

func Init() {
	inventory.Init()
	order.Init()
	product.Init()
	quote.Init()
}
