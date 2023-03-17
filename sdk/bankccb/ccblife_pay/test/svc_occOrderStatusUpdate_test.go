package test

import (
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"testing"
)

// 服务商订单更新
func TestOrderStatusUpdate(t *testing.T) {
	conf := getConfig()
	rs, err := ccblife_pay.OrderStatusUpdate(conf, ccblife_pay.OrderStatusUpdateParams{
		UserId:         "YSM202112210902914",
		OrderId:        "16301406188963962893",
		InformId:       ccblife_pay.INFORM_ID_0,
		OrderDt:        "20230227173944",
		TotalAmt:       "0.02",
		PayAmt:         "0.02",
		DiscountAmt:    "",
		PayStatus:      ccblife_pay.ORDER_STATUS_1,
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
	t.Log("ok.", rs.IsSuccess)
}
