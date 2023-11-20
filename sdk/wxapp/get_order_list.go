package wxapp

// GetOrderList 查询订单列表
func GetOrderList(accessToken string, params *GetOrderListParams) (list []*OrderItem, err error) {
	rs, err := post[GetOrderListResult]("https://api.weixin.qq.com/wxa/sec/order/get_order_list?access_token="+accessToken, params)
	if err != nil {
		return
	}
	list = rs.OrderList
	return
}

type GetOrderListParams struct {
	PayTimeRange *PayTimeRange `json:"pay_time_range"` //非必填， 支付时间所属范围。
	OrderState   OrderState    `json:"order_state"`    //订单状态枚举：(1) 待发货；(2) 已发货；(3) 确认收货；(4) 交易完成；(5) 已退款。
	Openid       string        `json:"openid"`         //
	PageSize     int64         `json:"page_size"`      // 默认100
}

type PayTimeRange struct {
	BeginTime int64 `json:"begin_time"` //起始时间，时间戳形式，不填则视为从0开始。
	EndTime   int64 `json:"end_time"`   //结束时间（含），时间戳形式，不填则视为32位无符号整型的最大值。
}

// OrderState 订单状态枚举
type OrderState int64

const (
	OrderState1 OrderState = 1
	OrderState2 OrderState = 2
	OrderState3 OrderState = 3
	OrderState4 OrderState = 4
	OrderState5 OrderState = 5
)

type GetOrderListResult struct {
	Errcode   int64        `json:"errcode"` // 0=成功，其它失败
	Errmsg    string       `json:"errmsg"`
	OrderList []*OrderItem `json:"order_list"` //支付单信息列表。
}

type OrderItem struct {
	TransactionId   string `json:"transaction_id"`    //原支付交易对应的微信订单号。
	MerchantId      string `json:"merchant_id"`       //支付下单商户的商户号，由微信支付生成并下发。
	MerchantTradeNo string `json:"merchant_trade_no"` //商户系统内部订单号，只能是数字、大小写字母`_-*`且在同一个商户号下唯一。
	Openid          string `json:"openid"`            //支付者openid。
}
