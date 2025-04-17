package controller

import "surgicalsteel-shop/service/quote"

var (
	getBuyerQuoteById quote.FnGetBuyerQuoteById
)

func Init() {
	getBuyerQuoteById = quote.GetBuyerQuoteById
}
