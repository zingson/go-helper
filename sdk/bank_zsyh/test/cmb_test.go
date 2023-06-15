package test

import (
	_ "embed"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/hid"
	"github.com/zingson/go-helper/htime"
	"github.com/zingson/go-helper/sdk/bank_zsyh"
	"os"
	"testing"
	"time"
)

// 生产环境
//
// //go:embed .secret/production.json

// 测试环境
//
//go:embed .secret/test.json
var configStr string

var conf *bank_zsyh.Config

func init() {
	err := json.Unmarshal([]byte(configStr), &conf)
	if err != nil {
		logrus.Error(err.Error())
	}
}

func init() {
	_ = os.MkdirAll("logs/"+time.Now().Format("200601"), 0600)
	file, err := os.OpenFile("logs/"+time.Now().Format("200601")+"/"+time.Now().Format("20060102T15")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	logrus.SetLevel(logrus.DebugLevel)
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
	resHtml, err := bank_zsyh.OneCardPayAPIForm(conf, bank_zsyh.OneCardPayAPIReq{
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
	orderNo := hid.G20()
	_, err := bank_zsyh.RefundAPI(conf, bank_zsyh.RefundAPIReq{
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
	orderNo := "2023021714592336af1d1d826e64e3d6"
	rspData, err := bank_zsyh.QuerySingleOrderAPI(conf, bank_zsyh.QuerySingleOrderAPIReq{
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
	orderNo := "2023021714592336af1d1d826e64e3d6"
	rspData, err := bank_zsyh.QuerySettledRefund(conf, bank_zsyh.QuerySettledRefundReq{
		DateTime:         htime.NowF14(),
		BranchNo:         conf.BranchNo,
		MerchantNo:       conf.MerchantNo,
		Type:             "B",
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
