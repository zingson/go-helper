package banknbcb

import (
	"root/hid"
	"testing"
)

/*
测试用：
易收宝商户号：33302125311054
*/
var cfg = &Config{
	ServiceUrl: "https://weixinsupport.nbcb.com.cn/qrcode1/payOut.do",
	//ServiceUrl: "https://ipay.nbcb.com.cn/qrcode/payOut.do",
	MchntCd:    "333021253111052",
	MchPriKey:  "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCZRt6obgml+2GMdVuWIYCBvbf0Sd7erelPS7NyBsUuSQ8Lk3hPtgiq/ulVaHF2hMFiTQPA02lSIvGKkvtCZrim812izewBsceG3ib9YaR8LjcZBuHF2dXJykL7xWWhW3ycmLcpim5QmohkbELw+vyI+UW7bWE/EYm7jYptMzzdgs6vVEy7ALkgfSR82DGR3ci/H1AW5Vu5A8LkHnY8Cp1sZUFGrS1c6VsiD59A4Jp4DKdvGFLpjaMnQBUzZWY+DzbtsQNZcAcJUJ2KMvwaQpqw9l7OdrBq/xQ7fs/fDVj51PUSzC84WtP9N9PkYq9X84p+syOdszKSfYRZVMGOELkTAgMBAAECggEAW1pV0mTi8z5EAYbgszX8dVcxkDOG1YkpiM1BgjBuzQtWIDwgdMG1oNSVFQZOuaid6YylNAPMvdt9wm2fuw+l9jsOD75Tbx8aIFO/QT00355b0Fq9rUILnV0jVxNrYyQ3vM14PRX4canoqJGtxBqS8MBAw1iJoSE+yb2uRy9Gk246xfFQuNb4ERjeLn0lZlUxeXD8SIEkdcIZ5YARlLmRcVloSFsv3FrWxZ9WGaQQWBPTCIhA/4Dt4LmyjiWwKg4BjZeswBm3w9OyoWFtqC9rrURQoyP2t1G+QHtk7mBZwNCARXTQZxblW5tGE+nzZCDz8XBPqjDncq1jIgzIURCKwQKBgQDIRFhbYfzdPYCSbZDod6sscB1tl+k/zlOLReYcBmgBqKiU+M7kFAwTEYxIuqXXBUTmxbL0qKTYOOokg8vYUCwX4TCIQBR5repzZsyJshvT8ZSuvqnw5ArUmUIgRrNtvZyVtUu6b4gREKj89KTcSfR5Q4IwSovclGKd4tzkt8ooCwKBgQDD7sxKanNnA9RH+i3N4+Mf7BqsgeKsUnYFE3jFdEwUyPzdlC7s7arLiP0CNP951SsJ6DbghdhR6CASI3/frDOiKHdlDYSbgOr1irDu3sWsQ8yO5uYWQTCqlKwsc2DQGA2vhWv4IK2O0Pl5vmBeYSbIwRns8BsHoq8j6rywyCpwGQKBgFb8aKx4qU6nkhsIADMZF36bzAx4OVX/loYd/E8b8T0XNvJOB/9FPeFic957Q+FbZowePxbJ0aAhSIJHxNjWKfDNXTkxLDOV+QnbEuiUNkYe3ofwRPxe8N1bHD/Mtc7q1wKn1pbKLv3KkLwevyT38nphekDjgFB1G0ic63lk4yf1AoGATXGl7gQFeUJaNdj9IdRhgcyg/m6YWeR+IaOAqQs+xzhqOmrH4X5PdAPBfY3VLSLE0DWo+zXsOvO1OXYupQo7nmzARIEDWTOrq0IWjVQgbeaehB1f9Ivv4HzDUQ6Jxba8MhRaiMjh0QFommZVUPi1i8RHqw98n5f+AgRWcMmtfIECgYEAirK2XEsGkAXKuFG8vWi2z5m0qMDtTlnGYW2LJJIDI4V1X7gy8EoE251wJS2v63BLXNuJ/OkSV0bQHFHCesRghYtm7BV/X1XE1rAHDMWSKwHsllJ1INmJ+Bv5QUa77PbrMVXfBxDDMPZVGT4bLhKcVh9JthJz/ihKZx2tyZTkGbs=",
	NbcbPubKey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmUbeqG4JpfthjHVbliGAgb239Ene3q3pT0uzcgbFLkkPC5N4T7YIqv7pVWhxdoTBYk0DwNNpUiLxipL7Qma4pvNdos3sAbHHht4m/WGkfC43GQbhxdnVycpC+8VloVt8nJi3KYpuUJqIZGxC8Pr8iPlFu21hPxGJu42KbTM83YLOr1RMuwC5IH0kfNgxkd3Ivx9QFuVbuQPC5B52PAqdbGVBRq0tXOlbIg+fQOCaeAynbxhS6Y2jJ0AVM2VmPg827bEDWXAHCVCdijL8GkKasPZeznawav8UO37P3w1Y+dT1EswvOFrT/TfT5GKvV/OKfrMjnbMykn2EWVTBjhC5EwIDAQAB",
}

// JSAPI下单
func TestJsPay(t *testing.T) {
	r, err := JsPay(cfg, &JsPayParams{
		ShopId:          "",
		SubAppid:        "wx309bfee2db775cdd",
		SubOpenid:       "oyozZ5CW2ifCzk6kTEzQRNfhOraI", //fxc zengs  openid oyozZ5CW2ifCzk6kTEzQRNfhOraI
		NotifyUrl:       "",
		TraceNo:         "11111111111111",
		TotalFee:        "1",
		IsCredit:        "",
		Remark:          "支付1分钱",
		DuplicateNoFlag: "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}

// 查询订单状态
func TestQueryOrderStatus(t *testing.T) {
	r, err := QueryOrderStatus(cfg, "14666642860036751360", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}

// 订单退款
func TestRefund(t *testing.T) {
	Refund(cfg, &RefundParams{
		StaffId:        "",
		RefundFee:      "1",
		TraceNo:        hid.G20(),
		OrigTraceNo:    "14680471552425574408",
		OrigOutTradeId: "20211207103811611188639490812677",
		Remark:         "退卡",
	})
}
