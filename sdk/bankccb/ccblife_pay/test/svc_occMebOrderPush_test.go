package test

import (
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"testing"
)

// 测试订单推送
func TestSVC_occMebOrderPush(t *testing.T) {
	conf := getConfig()
	ccblife_pay.MebOrderPush(conf, &ccblife_pay.OrderPushParams{
		UserId:        "",
		OrderId:       "",
		OrderDt:       "",
		TotalAmt:      "",
		PayAmt:        "",
		DiscountAmt:   "",
		OrderStatus:   "",
		RefundStatus:  "",
		MctNm:         "",
		CusOrderUrl:   "",
		OccMctLogoUrl: "",
	})
}
