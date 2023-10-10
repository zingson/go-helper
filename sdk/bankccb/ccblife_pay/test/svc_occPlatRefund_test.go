package test

import (
	"github.com/zingson/go-helper/sdk/bankccb/ccblife_pay"
	"testing"
)

// 退款测试
func TestRefund(t *testing.T) {
	conf := getConfig()
	cldBody, err := ccblife_pay.PlatRefund(conf, "17111484196507525138", 1, "20231009063549")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(cldBody.OrderNum)
	t.Log("refund success")
}
