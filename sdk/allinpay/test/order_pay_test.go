package test

import (
	allinpay "root/src/sdk/allinpay"
	"testing"
)

// 下单测试
func TestOrderPay(t *testing.T) {
	r, err := allinpay.Pay(config(), &allinpay.PayParams{
		Trxamt:        "1",
		Reqsn:         "11111111112211234",
		Paytype:       allinpay.PAY_TYPE_W06,
		Body:          "",
		Remark:        "",
		Validtime:     "",
		Acct:          "oSxaj4qF_DkjV2QU4GmS0FAaa6TU",
		NotifyUrl:     "",
		LimitPay:      "",
		SubAppid:      "wx4cf01a042bcc7599",
		GoodsTag:      "",
		Benefitdetail: "",
		FrontUrl:      "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)

}
