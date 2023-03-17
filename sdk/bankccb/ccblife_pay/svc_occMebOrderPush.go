package ccblife_pay

// MebOrderPush 服务方订单推送
// 仅仅是推送到建行生活APP进行订单展示
func MebOrderPush(conf *Config, cldBody OrderPushParams) (v OrderPushBody, err error) {
	return Call[OrderPushBody](conf, conf.ServiceSvcjson, "svc_occMebOrderPush", cldBody)
}

type OrderPushParams struct {
	UserId         string       `json:"USER_ID"`          //Y 客户编号，建行生活的会员编号
	OrderId        string       `json:"ORDER_ID"`         //Y 订单编号
	OrderDt        string       `json:"ORDER_DT"`         //Y 订单日期 yyyyMMddHHmmss
	TotalAmt       string       `json:"TOTAL_AMT"`        //Y 订单原金额, 单位：元
	PayAmt         string       `json:"PAY_AMT"`          //N 订单实际支付金额，单位：元  支付网关支付金额。此处如果为空必须在状态变更时推送
	DiscountAmt    string       `json:"DISCOUNT_AMT"`     //N 第三方平台优惠金额 第三方平台优惠金额。此处如果为空必须在状态变更时推送。
	OrderStatus    OrderStatus  `json:"ORDER_STATUS"`     //Y 订单状态 0-待支付  1-支付成功 2-已过期 3-支付失败  ,4-取消
	RefundStatus   RefundStatus `json:"REFUND_STATUS"`    //Y 退款状态 0-无退款  1-退款申请 2-已退款 3-部分退款
	InvDt          string       `json:"INV_DT"`           //N 订单过期日期 yyyyMMddHHmmss
	MctNm          string       `json:"MCT_NM"`           //Y 商户名称
	CusOrderUrl    string       `json:"CUS_ORDER_URL"`    //N 订单详情地址(需要推送完整的订单详情URL) ，非必填
	OccMctLogoUrl  string       `json:"OCC_MCT_LOGO_URL"` //N 服务方商户logo图片地址必须以不区分大小写的http://或者https://开头
	PayFlowId      string       `json:"PAY_FLOW_ID"`      //N 支付流水号
	PayUserId      string       `json:"PAY_USER_ID"`      //N 支付客户编号,非建行
	TotalRefundAmt string       `json:"TOTAL_REFUND_AMT"` //N 累计退款金额
	PlatOrderType  string       `json:"PLAT_ORDER_TYPE"`  //N 服务方订单类型  T0000-普通类型 T0001-洗车 T0002-加油 T0003-停车 T0004-修车 T0005-充电 T0006-年检代办 T0007-道路救援 T0008-云南中石油充值
}

// OrderStatus 订单状态 0-待支付  1-支付成功 2-已过期 3-支付失败  ,4-取消
type OrderStatus string

const (
	ORDER_STATUS_0 OrderStatus = "0"
	ORDER_STATUS_1 OrderStatus = "1"
	ORDER_STATUS_2 OrderStatus = "2"
	ORDER_STATUS_3 OrderStatus = "3"
	ORDER_STATUS_4 OrderStatus = "4"
)

var OrderStatusLabel = map[OrderStatus]string{
	ORDER_STATUS_0: "待支付",
	ORDER_STATUS_1: "支付成功",
	ORDER_STATUS_2: "已过期",
	ORDER_STATUS_3: "支付失败",
	ORDER_STATUS_4: "已取消",
}

// RefundStatus 退款状态 0-无退款  1-退款申请 2-已退款 3-部分退款
type RefundStatus string

const (
	REFUND_STATUS_0 RefundStatus = "0"
	REFUND_STATUS_1 RefundStatus = "1"
	REFUND_STATUS_2 RefundStatus = "2"
	REFUND_STATUS_3 RefundStatus = "3"
)

var RefundStatusLabel = map[RefundStatus]string{
	REFUND_STATUS_0: "无退款",
	REFUND_STATUS_1: "退款申请",
	REFUND_STATUS_2: "已退款",
	REFUND_STATUS_3: "部分退款",
}

type OrderPushBody struct {
	CcbDiscountAmt     string `json:"CCB_DISCOUNT_AMT"`      //建行支付侧优惠金额		15,2		C		N		在建行支付网关产生的优惠总金额
	CcbDiscountAmtDesc string `json:"CCB_DISCOUNT_AMT_DESC"` //建行支付侧优惠定义		1000		C		N		"建行支付侧优惠定义（各金额之和等于建行支付侧优惠金额） 格式：名称=金额|@|名称=金额 例如：优惠券A=0.05|@|活动A=0.99"
}
