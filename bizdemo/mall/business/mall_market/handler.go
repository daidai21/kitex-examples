package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_market/kitex_gen/cmp/ecom/market"
)

// MarketServiceImpl implements the last service interface defined in the IDL.
type MarketServiceImpl struct{}

// AddCoupon implements the MarketServiceImpl interface.
func (s *MarketServiceImpl) AddCoupon(ctx context.Context, req *market.AddCouponReq) (resp *market.AddCouponResp, err error) {
	// TODO: Your code here...
	return
}

// GetCouponById implements the MarketServiceImpl interface.
func (s *MarketServiceImpl) GetCouponById(ctx context.Context, req *market.GetCouponByIdReq) (resp *market.GetCouponByIdResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateCoupon implements the MarketServiceImpl interface.
func (s *MarketServiceImpl) UpdateCoupon(ctx context.Context, req *market.UpdateCouponReq) (resp *market.UpdateCouponResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteCoupon implements the MarketServiceImpl interface.
func (s *MarketServiceImpl) DeleteCoupon(ctx context.Context, req *market.DeleteCouponReq) (resp *market.DeleteCouponResp, err error) {
	// TODO: Your code here...
	return
}

// UsingCoupon implements the MarketServiceImpl interface.
func (s *MarketServiceImpl) UsingCoupon(ctx context.Context, req *market.UsingCouponReq) (resp *market.UsingCouponResp, err error) {
	// TODO: Your code here...
	return
}

// RetCoupon implements the MarketServiceImpl interface.
func (s *MarketServiceImpl) RetCoupon(ctx context.Context, req *market.RetCouponReq) (resp *market.RetCouponResp, err error) {
	// TODO: Your code here...
	return
}
