package cmbnetpay

import (
	"encoding/json"
)

//CardPay 一网通支付 接口文档:http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/OneCardPayAPI
//生产环境 https://netpay.cmbchina.com/netpayment/BaseHttp.dll?MB_EUserPay
//测试环境 http://paytest.cmburl.cn:801/netpayment/BaseHttp.dll?MB_EUserPay
func CardPay(conf *Config, req *CardPayReq) (resHtml string, err error) {
	reqMap := StructToMap(req.ReqData)
	waitForSignStr := SortMap(reqMap, true) + "&" + conf.Merkey
	req.Sign = Sha256Sign(waitForSignStr)

	rBytes, err := json.Marshal(req)
	if err != nil {
		return
	}
	resHtml, err = PostForm(conf, conf.ApiUrl, string(rBytes))
	if err != nil {
		return
	}
	return resHtml, nil
}

type CardPayReq struct {
	FixedParams
	ReqData *PayRequestData `json:"reqData"`
}

type PayRequestData struct {
	DateTime       string `json:"dateTime"`       //请求时间,格式：yyyyMMddHHmmss 商户发起该请求的当前时间，精确到秒
	BranchNo       string `json:"branchNo"`       //分行号，4位数字
	MerchantNo     string `json:"merchantNo"`     //商户号，6位数字
	Date           string `json:"date"`           //订单日期,格式：yyyyMMdd
	OrderNo        string `json:"orderNo"`        //订单号。 支持6-32位（含6和32）间任意位数的订单号，支持不固定位数，支持数字+字母（大小字母）随意组合。由商户生成，一天内不能重复。 订单日期+订单号唯一定位一笔订单。
	Amount         string `json:"amount"`         //订单金额，格式：xxxx.xx 固定两位小数，最大11位整数
	ExpireTimeSpan string `json:"expireTimeSpan"` //过期时间跨度，必须为大于零的整数，单位为分钟。该参数指定当前支付请求必须在指定时间跨度内完成（从系统收到支付请求开始计时），否则按过期处理。一般适用于航空客票等对交易完成时间敏感的支付请求。
	PayNoticeUrl   string `json:"payNoticeUrl"`   //商户接收成功支付结果通知的地址。
	//PayNoticePara  string `json:"payNoticePara"`  //成功支付结果通知附加参数,该参数在发送成功支付结果通知时，将原样返回商户注意：该参数可为空，商户如果需要不止一个参数，可以自行把参数组合、拼装，但组合后的结果不能带有’&’字符。
	//ReturnUrl            string `json:"returnUrl"`
	//ClientIP             string `json:"clientIP"`
	//CardType             string `json:"cardType"`
	//AgrNo                string `json:"agrNo"`
	//MerchantSerialNo     string `json:"merchantSerialNo"`
	//UserID               string `json:"userID"`
	//Mobile               string `json:"mobile"`
	//Lon                  string `json:"lon"`
	//Lat                  string `json:"lat"`
	//RiskLevel            string `json:"riskLevel"`
	//SignNoticeUrl        string `json:"signNoticeUrl"`
	//SignNoticePara       string `json:"signNoticePara"`
	//ExtendInfo           string `json:"extendInfo"`
	//ExtendInfoEncrypType string `json:"extendInfoEncrypType"`
}
