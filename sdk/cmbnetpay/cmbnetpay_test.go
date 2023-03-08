package cmbnetpay

import (
	"github.com/sirupsen/logrus"
	"github.com/zingson/goh/hid"
	"github.com/zingson/goh/htime"
	"testing"
)

// TestQueryFbPubKey 查询招行公钥
func TestQueryFbPubKey(t *testing.T) {
	conf := &Config{
		BranchNo:   "0755",
		MerchantNo: "058624",
		Merkey:     "1234567890abcABC",
		ApiUrl:     QUERY_FBPUBKEY_URL_TEST,
	}
	pubkey, err := QueryFbPubKey(conf)
	if err != nil {
		logrus.Infof("TestQueryFbPubKey, error:%s", err.Error())
	}
	logrus.Infof("TestQueryFbPubKey, pubkey:%s", pubkey)
}

// TestCardPay 支付
func TestCardPay(t *testing.T) {
	conf := &Config{
		BranchNo:   "0755",
		MerchantNo: "058624",
		Merkey:     "1234567890abcABC",
		ApiUrl:     CARD_PAY_URL_TEST,
	}

	orderNo := hid.G20()
	requestData := &PayRequestData{
		DateTime:       htime.NowF14(),
		BranchNo:       conf.BranchNo,
		MerchantNo:     conf.MerchantNo,
		Date:           htime.NowF8(),
		OrderNo:        orderNo,
		Amount:         "0.01",
		ExpireTimeSpan: "15",
		//PayNoticeUrl:   config.HimktC().BaseUrl + "/gw/" + config.Name + "/pay/cmbnetpay/notify",
		//PayNoticePara:  orderNo + "|" + conf.MerchantNo,
	}
	resHtml, err := CardPay(conf, &CardPayReq{
		FixedParams: GetFixedParams(),
		ReqData:     requestData,
	})
	if err != nil {
		logrus.Infof("TestCardPay, error:%s", err.Error())
	}
	logrus.Infof("TestCardPay, response:%s", resHtml)
	//ctx.Data(200, "text/html", []byte(resHtml))
}

// TestRefund 退款
func TestRefund(t *testing.T) {
	conf := &Config{
		BranchNo:   "0755",
		MerchantNo: "058624",
		Merkey:     "1234567890abcABC",
		ApiUrl:     REFUND_URL_TEST,
	}

	orderNo := hid.G20()
	requestData := &RefundReqData{
		DateTime:       htime.NowF14(),
		BranchNo:       conf.BranchNo,
		MerchantNo:     conf.MerchantNo,
		Date:           htime.NowF8(),
		OrderNo:        orderNo,
		RefundSerialNo: hid.G20(),
		Amount:         "0.01",
	}
	res, err := Refund(conf, &RefundReq{
		FixedParams: GetFixedParams(),
		ReqData:     requestData,
	})
	if err != nil {
		logrus.Infof("TestRefund, error:%s", err.Error())
	}
	logrus.Infof("TestRefund, response:%v", res)
}

// TestQueryOrder 单笔订单查询
func TestQueryOrder(t *testing.T) {
	conf := &Config{
		BranchNo:   "0755",
		MerchantNo: "058613",
		Merkey:     "1234567890abcABC",
		ApiUrl:     QUERY_ORDER_URL_TEST,
	}

	orderNo := "2023021714592336af1d1d826e64e3d6"
	requestData := &QueryOrderReqData{
		DateTime:   htime.NowF14(),
		BranchNo:   conf.BranchNo,
		MerchantNo: conf.MerchantNo,
		Date:       htime.NowF8(),
		OrderNo:    orderNo,
		Type:       "B",
	}
	res, err := QueryOrder(conf, &QueryOrderReq{
		FixedParams: GetFixedParams(),
		ReqData:     requestData,
	})
	if err != nil {
		logrus.Infof("TestQueryOrder, error:%s", err.Error())
	}
	logrus.Infof("TestQueryOrder, response:%v", res)
}
