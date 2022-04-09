package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/pack"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/service"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
)

// ShopServiceImpl implements the last service interface defined in the IDL.
type ShopServiceImpl struct{}

// SettleShop implements the ShopServiceImpl interface.
func (s *ShopServiceImpl) SettleShop(ctx context.Context, req *shop.SettleShopReq) (resp *shop.SettleShopResp, err error) {
	resp = shop.NewSettleShopResp()
	if req.ShopName == "" {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	shopId, err := service.NewShopService(ctx).SettleShop(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.ShopId = shopId
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetShopIdByUserId implements the ShopServiceImpl interface.
func (s *ShopServiceImpl) GetShopIdByUserId(ctx context.Context, req *shop.GetShopIdByUserIdReq) (resp *shop.GetShopIdByUserIdResp, err error) {
	resp = shop.NewGetShopIdByUserIdResp()
	shopId, err := service.NewShopService(ctx).GetShopIdByUserId(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.ShopId = shopId
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
