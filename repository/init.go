package repository

import (
	"surgicalsteel-shop/repository/base"
	"surgicalsteel-shop/repository/inventory"
	"surgicalsteel-shop/repository/order"
	"surgicalsteel-shop/repository/order_item"
	"surgicalsteel-shop/repository/product"
	"surgicalsteel-shop/repository/quote"
)

func Init() {
	inventory.Init()
	order.Init()
	product.Init()
	quote.Init()
	base.Init()
	order_item.Init()
}
