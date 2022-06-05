package service

import "context"

type MarketService struct {
	ctx context.Context
}

func NewMarketService(ctx context.Context) *MarketService {
	return &MarketService{ctx: ctx}
}

// TODO
