package ccblife_pay

import (
	"bytes"
	_ "embed"
	"html/template"
)

//go:embed svc_occOrderDetail.html
var htmlTpl string

var htmlTemplate *template.Template

// OrderDetailHtml 订单详情查看
// 提供查看订单的默认页面
func OrderDetailHtml(data OrderDetailParam) (html string, err error) {

	if data.StatusLabel == "" {
		if data.RefundStatus != REFUND_STATUS_0 {
			data.StatusLabel = RefundStatusLabel[data.RefundStatus]
		} else {
			data.StatusLabel = OrderStatusLabel[data.OrderStatus]
		}
	}

	if htmlTemplate == nil {
		htmlTemplate, err = template.New("svc_occOrderDetail").Parse(htmlTpl)
		if err != nil {
			return
		}
	}

	var buf bytes.Buffer
	err = htmlTemplate.Execute(&buf, data)
	if err != nil {
		return
	}
	html = buf.String()
	return
}

type OrderDetailParam struct {
	OrderId        string       //Y 订单编号
	OrderDt        string       //Y 订单日期 yyyy-MM-dd HH:mm:ss
	TotalAmt       string       //Y 订单原金额, 单位：元
	PayAmt         string       //N 订单实际支付金额，单位：元  支付网关支付金额。此处如果为空必须在状态变更时推送
	DiscountAmt    string       //N 第三方平台优惠金额 第三方平台优惠金额。此处如果为空必须在状态变更时推送。 单位：元
	GoodsImg       string       //Y 商品图链接
	GoodsName      string       //Y 商品名称
	GoodsNum       string       //Y 商品数量
	StatusLabel    string       //Y 订单状态文字
	OrderStatus    OrderStatus  //Y 订单状态 0-待支付  1-支付成功 2-已过期 3-支付失败  ,4-取消
	RefundStatus   RefundStatus //Y 退款状态 0-无退款  1-退款申请 2-已退款 3-部分退款
	InvDt          string       //N 订单过期日期 yyyyMMddHHmmss
	MctNm          string       //Y 商户名称
	OccMctLogoUrl  string       //N 服务方商户logo图片地址必须以不区分大小写的http://或者https://开头
	PayFlowId      string       //N 支付流水号
	PayUserId      string       //N 支付客户编号,非建行
	TotalRefundAmt string       //N 累计退款金额
	PlatOrderType  string       //N 服务方订单类型  T0000-普通类型 T0001-洗车 T0002-加油 T0003-停车 T0004-修车 T0005-充电 T0006-年检代办 T0007-道路救援 T0008-云南中石油充值
}
