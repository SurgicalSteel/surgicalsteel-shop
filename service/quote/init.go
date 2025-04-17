package quote

import "surgicalsteel-shop/repository/quote"

var (
	getQuoteById quote.FnGetQuoteById
)

func Init() {
	getQuoteById = quote.GetQuoteById

}
