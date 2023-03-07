package v3

type PayService struct {
	*Client
}

func (c *Client) Pay() *PayService {
	return &PayService{Client: c}
}

//TradeType 交易类型
type TradeType string

const (
	JSAPI    TradeType = "JSAPI"    //公众号支付/小程序
	NATIVE   TradeType = "NATIVE"   //扫码支付
	APP      TradeType = "APP"      //APP支付
	MICROPAY TradeType = "MICROPAY" //付款码支付
	MWEB     TradeType = "MWEB"     //H5支付
	FACEPAY  TradeType = "FACEPAY"  //刷脸支付
)

/*
交易状态，枚举值：
SUCCESS：支付成功
REFUND：转入退款
NOTPAY：未支付
CLOSED：已关闭
REVOKED：已撤销（付款码支付）
USERPAYING：用户支付中（付款码支付）
PAYERROR：支付失败(其他原因，如银行返回失败)
*/
type TradeState string

var (
	TRADE_STATE_SUCCESS    TradeState = "SUCCESS"    //支付成功
	TRADE_STATE_REFUND     TradeState = "REFUND"     //转入退款
	TRADE_STATE_NOTPAY     TradeState = "NOTPAY"     //未支付
	TRADE_STATE_CLOSED     TradeState = "CLOSED"     //已关闭
	TRADE_STATE_REVOKED    TradeState = "REVOKED"    //已撤销（付款码支付）
	TRADE_STATE_USERPAYING TradeState = "USERPAYING" //用户支付中（付款码支付）
	TRADE_STATE_PAYERROR   TradeState = "PAYERROR"   //支付失败(其他原因，如银行返回失败)
)
