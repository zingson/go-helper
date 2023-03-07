package banknbcb

// 退款状态
type RefundStatus string

const (
	RS02 RefundStatus = "02" // 退款成功
	RS11 RefundStatus = "11" // 退款中
	RS12 RefundStatus = "12" // 退款错误
)
