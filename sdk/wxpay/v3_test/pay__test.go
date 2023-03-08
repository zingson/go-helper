package v3_test

import (
	"github.com/zingson/go-helper/sdk/wxpay/v3"
	"testing"
)

// JSAPI 支付
func TestPayJsapi(t *testing.T) {
	pay := Client().Pay()
	p, err := pay.Jsapi(&v3.JsapiPay{
		Appid:       "wx239c521c61221a8a",
		Mchid:       pay.Mchid,
		Description: "描述",
		OutTradeNo:  "A123456789011",
		TimeExpire:  "",
		Attach:      "1",
		NotifyUrl:   "https://msd.himkt.cn/gw/xxx",
		GoodsTag:    "",
		Amount:      &v3.PayAmount{Total: 1},
		Payer:       &v3.PayPayer{Openid: "oW-G-0hraaeGSfxrx_q9AfYRev60"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(p.JSON())
}
