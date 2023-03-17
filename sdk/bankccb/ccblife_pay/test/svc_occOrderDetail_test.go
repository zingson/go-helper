package test

import (
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"testing"
)

func TestOrderDetailHtml(t *testing.T) {
	html, err := ccblife_pay.OrderDetailHtml(ccblife_pay.OrderDetailParam{
		OrderId:        "",
		OrderDt:        "",
		TotalAmt:       "",
		PayAmt:         "",
		DiscountAmt:    "",
		GoodsImg:       "",
		GoodsName:      "",
		GoodsNum:       "",
		StatusLabel:    "",
		OrderStatus:    "",
		RefundStatus:   "",
		InvDt:          "",
		MctNm:          "",
		OccMctLogoUrl:  "",
		PayFlowId:      "",
		PayUserId:      "",
		TotalRefundAmt: "",
		PlatOrderType:  "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(html)
}
