package v3

import (
	"fmt"
)

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_6.shtml

//CouponInfo 券详情查询
func (mkt *MktService) CouponInfo(params CouponInfoParams) (result *CouponInfoResult, err error) {
	err = mkt.HttpGet(fmt.Sprintf("/v3/marketing/favor/users/%s/coupons/%s?appid=%s", params.Openid, params.CouponId, params.Appid), &result)
	return
}

type CouponInfoParams struct {
	Appid    string `json:"appid"`     // 授权appid ，需要与发券商户绑定
	Openid   string `json:"openid"`    // 授权openid
	CouponId string `json:"coupon_id"` // 券ID，发券接口返回
}

type CouponInfoResult struct {
}
