package ccblife_pay

// 服务方订单推送
func SVC_occMebOrderPush(conf *Config, args *OrderPushParams) (err error) {
	_, err = Post(conf, conf.ServiceSvcjson, "svc_occMebOrderPush", args)
	if err != nil {
		return
	}
	return
}

type OrderPushParams struct {
	USER_ID          string `json:"USER_ID"`
	ORDER_ID         string `json:"ORDER_ID"`
	ORDER_DT         string `json:"ORDER_DT"`
	TOTAL_AMT        string `json:"TOTAL_AMT"`
	PAY_AMT          string `json:"PAY_AMT"`
	DISCOUNT_AMT     string `json:"DISCOUNT_AMT"`
	ORDER_STATUS     string `json:"ORDER_STATUS"`     //0-待支付  1-支付成功 2-已过期 3-支付失败  ,4-取消
	REFUND_STATUS    string `json:"REFUND_STATUS"`    //0-无退款  1-退款申请 2-已退款 3-部分退款
	MCT_NM           string `json:"MCT_NM"`           //商户名称
	CUS_ORDER_URL    string `json:"CUS_ORDER_URL"`    // 订单详情地址(需要推送完整的订单详情URL) ，非必填
	OCC_MCT_LOGO_URL string `json:"OCC_MCT_LOGO_URL"` //服务方商户logo图片地址必须以不区分大小写的http://或者https://开头
}
