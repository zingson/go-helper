package bank_zhywt

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// OneCardPayAPI 一网通支付API  文档:http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/OneCardPayAPI
// 生产环境 https://netpay.cmbchina.com/netpayment/BaseHttp.dll?MB_EUserPay
// 测试环境 http://paytest.cmburl.cn:801/netpayment/BaseHttp.dll?MB_EUserPay
func OneCardPayAPI(conf *Config, reqData OneCardPayAPIReq) (rspData OneCardPayAPIRes, err error) {
	jsonRequestData := jsonMarshal(&RequestBody[OneCardPayAPIReq]{
		Version:  Version,
		Charset:  Charset,
		Sign:     Sha256Sign(SortMap(StructToMap(reqData), true) + "&" + conf.Merkey),
		SignType: SignType,
		ReqData:  reqData,
	})

	logrus.WithField("mchid", conf.BranchNo+conf.MerchantNo).Infof("招行一网通H5支付 下单 \n接口：%s  \n请求：%s  \n响应：%s", conf.NetpaymentUrl+"/netpayment/BaseHttp.dll?MB_EUserPay", jsonRequestData, "html")

	rspData = OneCardPayAPIRes{
		JsonRequestData: jsonRequestData,
		Charset:         Charset,
	}
	return
}

// OneCardPayAPIForm 返回表单HTML，前端通过表单跳转支付
func OneCardPayAPIForm(conf *Config, reqData OneCardPayAPIReq) (resHtml string, err error) {
	rspData, err := OneCardPayAPI(conf, reqData)
	if err != nil {
		return
	}
	resHtml = fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<form action="%s" method="POST" id="form1">
    <input type="hidden" name="jsonRequestData" value='%s'/>
    <input type="hidden" name="charset" value='%s'/>
</form>
<script>window.onload = function(){document.getElementById("form1").submit()}</script>
</body>
</html>`, conf.NetpaymentUrl+"/netpayment/BaseHttp.dll?MB_EUserPay", rspData.JsonRequestData, rspData.Charset)
	return
}

type OneCardPayAPIReq struct {
	DateTime             string `json:"dateTime"`             //M 请求时间,格式：yyyyMMddHHmmss 商户发起该请求的当前时间，精确到秒
	BranchNo             string `json:"branchNo"`             //M 分行号，4位数字
	MerchantNo           string `json:"merchantNo"`           //M 商户号，6位数字
	Date                 string `json:"date"`                 //M 订单日期,格式：yyyyMMdd
	OrderNo              string `json:"orderNo"`              //M 订单号。 支持6-32位（含6和32）间任意位数的订单号，支持不固定位数，支持数字+字母（大小字母）随意组合。由商户生成，一天内不能重复。 订单日期+订单号唯一定位一笔订单。
	Amount               string `json:"amount"`               //M 订单金额，单位元 格式：xxxx.xx 固定两位小数，最大11位整数
	ExpireTimeSpan       string `json:"expireTimeSpan"`       //M 过期时间跨度，必须为大于零的整数，单位为分钟。该参数指定当前支付请求必须在指定时间跨度内完成（从系统收到支付请求开始计时），否则按过期处理。一般适用于航空客票等对交易完成时间敏感的支付请求。
	PayNoticeUrl         string `json:"payNoticeUrl"`         //M 商户接收成功支付结果通知的地址。
	PayNoticePara        string `json:"payNoticePara"`        //O成功支付结果通知附加参数,该参数在发送成功支付结果通知时，将原样返回商户注意：该参数可为空，商户如果需要不止一个参数，可以自行把参数组合、拼装，但组合后的结果不能带有’&’字符。
	ReturnUrl            string `json:"returnUrl"`            //O
	ClientIP             string `json:"clientIP"`             //O
	CardType             string `json:"cardType"`             //O
	AgrNo                string `json:"agrNo"`                //O
	MerchantSerialNo     string `json:"merchantSerialNo"`     //C
	UserID               string `json:"userID"`               //O
	Mobile               string `json:"mobile"`               //O
	Lon                  string `json:"lon"`                  //O
	Lat                  string `json:"lat"`                  //O
	RiskLevel            string `json:"riskLevel"`            //O
	SignNoticeUrl        string `json:"signNoticeUrl"`        //O
	SignNoticePara       string `json:"signNoticePara"`       //O
	ExtendInfo           string `json:"extendInfo"`           //O
	ExtendInfoEncrypType string `json:"extendInfoEncrypType"` //O
}

type OneCardPayAPIRes struct {
	JsonRequestData string `json:"jsonRequestData"`
	Charset         string `json:"charset"`
}
