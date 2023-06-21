package test

import (
	_ "embed"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/hid"
	"github.com/zingson/go-helper/htime"
	"github.com/zingson/go-helper/sdk/bank_zhywt"
	"os"
	"testing"
)

// 生产环境
//
//go:embed .secret/production.json

// 测试环境
//
// //go:embed .secret/test.json
var configStr string

var conf *bank_zhywt.Config

func init() {
	err := json.Unmarshal([]byte(configStr), &conf)
	if err != nil {
		logrus.Error(err.Error())
	}
}

// TestQueryKeyAPI 查询招行公钥
func TestQueryKeyAPI(t *testing.T) {
	fb, err := conf.GetFbPubKey()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(fb)
}

// TestOneCardPayAPI 支付
func TestOneCardPayAPI(t *testing.T) {
	orderNo := hid.G20()
	resHtml, err := bank_zhywt.OneCardPayAPIForm(conf, bank_zhywt.OneCardPayAPIReq{
		DateTime:             htime.NowF14(),
		BranchNo:             conf.BranchNo,
		MerchantNo:           conf.MerchantNo,
		Date:                 htime.NowF8(),
		OrderNo:              orderNo,
		Amount:               "0.01",
		ExpireTimeSpan:       "15",
		PayNoticeUrl:         "https://msd.himkt.cn/gw/upmp/ywt/xx",
		PayNoticePara:        "",
		ReturnUrl:            "",
		ClientIP:             "",
		CardType:             "",
		AgrNo:                "",
		MerchantSerialNo:     "",
		UserID:               "",
		Mobile:               "",
		Lon:                  "",
		Lat:                  "",
		RiskLevel:            "",
		SignNoticeUrl:        "",
		SignNoticePara:       "",
		ExtendInfo:           "",
		ExtendInfoEncrypType: "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resHtml)

	_ = os.WriteFile("OneCardPayAPI.html", []byte(resHtml), 0600)
}

// TestRefundAPI 退款
func TestRefundAPI(t *testing.T) {
	orderNo := "16696016991384371202"
	_, err := bank_zhywt.RefundAPI(conf, bank_zhywt.RefundAPIReq{
		DateTime:       htime.NowF14(),
		BranchNo:       conf.BranchNo,
		MerchantNo:     conf.MerchantNo,
		Date:           htime.NowF8(),
		OrderNo:        orderNo,
		RefundSerialNo: hid.G20(),
		Amount:         "0.01",
		Desc:           "",
	})
	if err != nil {
		t.Error(err)
		return
	}
}

// TestQuerySingleOrderAPI 单笔订单查询
func TestQuerySingleOrderAPI(t *testing.T) {
	orderNo := "16696016991384371202"
	rspData, err := bank_zhywt.QuerySingleOrderAPI(conf, bank_zhywt.QuerySingleOrderAPIReq{
		DateTime:   htime.NowF14(),
		BranchNo:   conf.BranchNo,
		MerchantNo: conf.MerchantNo,
		Date:       htime.NowF8(),
		OrderNo:    orderNo,
		Type:       "B",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rspData)
}

// TestQuerySettledRefund 单笔退款查询
func TestQuerySettledRefund(t *testing.T) {
	orderNo := "16696016991384371202"
	rspData, err := bank_zhywt.QuerySettledRefund(conf, bank_zhywt.QuerySettledRefundReq{
		DateTime:         htime.NowF14(),
		BranchNo:         conf.BranchNo,
		MerchantNo:       conf.MerchantNo,
		Type:             "C",
		BankSerialNo:     "",
		Date:             htime.NowF8(),
		OrderNo:          orderNo,
		MerchantSerialNo: orderNo,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rspData)
}

// TestSuccessPayAPI 支付成功通知
func TestSuccessPayAPI(t *testing.T) {
	noticeBody := `{"charset":"UTF-8","noticeData":{"dateTime":"20230616193916","date":"20230616","amount":"0.01","bankDate":"20230616","orderNo":"16696709751232512016","cardType":"03","discountAmount":"0.00","noticeType":"BKPAYRTN","httpMethod":"POST","noticeSerialNo":"2023061616696709751232512016","merchantPara":"","discountFlag":"N","bankSerialNo":"226LZ0DO054AA100000A","noticeUrl":"https://msd.himkt.cn/gw/upmp/hc-upmp-order-ywth5pay-notify/0574630317","branchNo":"0574","merchantNo":"630317"},"sign":"scW9VguruCaII9+a8yhwohkZN1HPAx9ZEHQNpDa1XaUS+6iwPmp4fIMPwzbMmRym6AT5iPmWNxQkuAIoxObHqyshuKCs3R1sgmVxt2KDtCfasslwYjkIQNqbbcUkw0uS7QSXwuJkuZe2Q/jow/AvlEIbRS2r3XYGszzr7lTdcv8=","signType":"RSA","version":"1.0"}`
	noticeData, err := bank_zhywt.SuccessPayApi(conf, "/notice/path", []byte(noticeBody))
	if err != nil {
		t.Error(err.Error())
		return
	}
	b, _ := json.Marshal(noticeData)
	t.Log(string(b))
}
