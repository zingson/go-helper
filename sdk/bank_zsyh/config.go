package bank_zsyh

// Config 招行一网通接口配置参数
type Config struct {
	MerchservUrl  string `json:"merchserv_url"`  //https://merchserv.netpay.cmbchina.com/
	CmbBankB2BUrl string `json:"cmbbankb2b_url"` //https://b2b.cmbchina.com/
	NetpaymentUrl string `json:"netpayment_url"` //https://netpay.cmbchina.com/

	ServiceUrl string `json:"service_url"` //服务地址
	BranchNo   string `json:"branch_no"`   //商户分行号，4位数字	0755
	MerchantNo string `json:"merchant_no"` //商户号，6位数字	058624
	Merkey     string `json:"mer_key"`     //商户秘钥
	FbPubKey   string `json:"fb_pub_key"`  //用Base64编码的招行公钥
}

const (
	QUERY_FBPUBKEY_URL_PROD string = "https://b2b.cmbchina.com/CmbBank_B2B/UI/NetPay/DoBusiness.ashx"          //查询招行公钥APIURL产线地址
	QUERY_FBPUBKEY_URL_TEST string = "https://cmbchinab2b.bas.cmburl.cn/CmbBank_B2B/UI/NetPay/DoBusiness.ashx" //查询招行公钥APIURL测试地址

	CARD_PAY_URL_PROD string = "https://netpay.cmbchina.com/netpayment/BaseHttp.dll?MB_EUserPay"  //一网通支付APIURL产线地址
	CARD_PAY_URL_TEST string = "http://paytest.cmburl.cn:801/netpayment/BaseHttp.dll?MB_EUserPay" //一网通支付APIURL测试地址

	REFUND_URL_PROD string = "https://merchserv.netpay.cmbchina.com/merchserv/BaseHttp.dll?DoRefundV2" //一网通退款APIURL产线地址
	REFUND_URL_TEST string = "http://merchserv.cmburl.cn/merchserv/BaseHttp.dll?DoRefundV2"            //一网通退款APIURL测试地址

	QUERY_ORDER_URL_PROD string = "https://merchserv.netpay.cmbchina.com/merchserv/BaseHttp.dll?QuerySingleOrder" //一网通查询单笔订单APIURL产线地址
	QUERY_ORDER_URL_TEST string = "http://merchserv.cmburl.cn/merchserv/BaseHttp.dll?QuerySingleOrder"            //一网通查询单笔订单APIURL测试地址
)
