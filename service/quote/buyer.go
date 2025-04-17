package quote

import (
	"context"
	"errors"
	"surgicalsteel-shop/repository/quote"
)

type FnGetBuyerQuoteById func(ctx context.Context, id int64) (quote.SongQuote, error)

func GetBuyerQuoteById(ctx context.Context, id int64) (quote.SongQuote, error) {
	if id <= 0 {
		return quote.SongQuote{}, errors.New("invalid id")
	}

	quote, err := getQuoteById(ctx, id)
	if err != nil {
		return quote, err
	}
	if quote.Id == 0 {
		return quote, errors.New("quote not found")
	}
	return quote, err
}
