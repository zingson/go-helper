package test

import (
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"testing"
)

// 测试订单推送
func TestSVC_occMebOrderPush(t *testing.T) {
	conf := getConfig()
	_, err := ccblife_pay.MebOrderPush(conf, ccblife_pay.OrderPushParams{
		UserId:         "YSM202112210902914",
		OrderId:        "16301406188963962893",
		OrderDt:        "20230227173944",
		TotalAmt:       "0.02",
		PayAmt:         "",
		DiscountAmt:    "",
		OrderStatus:    ccblife_pay.ORDER_STATUS_1,
		RefundStatus:   ccblife_pay.REFUND_STATUS_0,
		InvDt:          "",
		MctNm:          "海脉云",
		CusOrderUrl:    "",
		OccMctLogoUrl:  "",
		PayFlowId:      "",
		PayUserId:      "",
		TotalRefundAmt: "",
		PlatOrderType:  "T0000",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("ok.")
}
