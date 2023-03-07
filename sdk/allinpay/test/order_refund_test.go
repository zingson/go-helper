package test

import (
	allinpay "root/src/sdk/allinpay"
	"testing"
)

// 退款交易测试
func TestOrderRefund(t *testing.T) {
	r, err := allinpay.Refund(config(), &allinpay.RefundParams{
		Trxamt:        "1",
		Reqsn:         "12345678",
		Oldreqsn:      "111111111122",
		Oldtrxid:      "112181560000967098",
		Remark:        "",
		Benefitdetail: "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}
