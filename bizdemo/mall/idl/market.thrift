// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

include "base.thrift"
namespace go cmp.ecom.market


// 优惠券使用范围： 品牌、店铺、平台通用、商品劵、类目
enum CouponType {
    Brand,
    Shop,
    Platform,
    Product,
    Category,
}

struct Coupon {
    1: i64 coupon_id,
    2: i64 use_scope_id,
    3: CouponType coupon_type,
    4: i64 start_time,
    5: i64 end_time,
    6: i64 total_stock,
    7: i64 real_stock,
    8: string coupon_name,
    9: i64 limit_low_price,
    10: i64  reduce_price,
}

struct AddCouponReq {
    2: i64 use_scope_id,
    3: CouponType coupon_type,
    4: i64 start_time,
    5: i64 end_time,
    6: i64 total_stock,
    7: i64 real_stock,
    8: string coupon_name,
    9: i64 limit_low_price,
    10: i64  reduce_price,
}


struct AddCouponResp {
    1: i64 coupon_id,

    255: base.BaseResp BaseResp
}

struct GetCouponByIdReq {
    1: i64 coupon_id,
}


struct GetCouponByIdResp {
    2: i64 use_scope_id,
    3: CouponType coupon_type,
    4: i64 start_time,
    5: i64 end_time,
    6: i64 total_stock,
    7: i64 real_stock,
    8: string coupon_name,
    9: i64 limit_low_price,
    10: i64  reduce_price,

    255: base.BaseResp BaseResp
}

struct UpdateCouponReq {
    1: i64 coupon_id,
    2: i64 use_scope_id,
    3: CouponType coupon_type,
    4: i64 start_time,
    5: i64 end_time,
    6: i64 total_stock,
    7: i64 real_stock,
    8: string coupon_name,
    9: i64 limit_low_price,
    10: i64  reduce_price
}


struct UpdateCouponResp {
    255: base.BaseResp BaseResp
}


struct DeleteCouponReq {
    1: i64 coupon_id
}


struct DeleteCouponResp {
    255: base.BaseResp BaseResp
}


struct UsingCouponReq {
    1: i64 product_id,
    2: i64 coupon_id,

    255: base.BaseResp BaseResp
}

struct UsingCouponResp {
    1: i64 reduce_price,

    255: base.BaseResp BaseResp
}

struct RetCouponReq {
    1: i64 coupon_id,

    255: base.BaseResp BaseResp
}

struct RetCouponResp {
    255: base.BaseResp BaseResp
}

service MarketService {
    // CRUD
    AddCouponResp AddCoupon(1: AddCouponReq req)
    GetCouponByIdResp GetCouponById(1: GetCouponByIdReq req)
    UpdateCouponResp UpdateCoupon(1: UpdateCouponReq req)
    DeleteCouponResp DeleteCoupon(1: DeleteCouponReq req)

    // TODO
    // 查询商品可用的劵列表
    // 没有考虑劵的是否能叠加和排他规则使用

    // 输入商品和劵的id，返回可以优惠的价格
    UsingCouponResp UsingCoupon(1: UsingCouponReq req)
    // 劵的库存 增加/回补
    RetCouponResp RetCoupon(1: RetCouponReq req)
}
