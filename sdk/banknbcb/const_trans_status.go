package banknbcb

// TransStatus 交易状态
type TransStatus string

const (
	TS00 TransStatus = "00" //"00"-- 交易初始
	TS01 TransStatus = "01" //"01"-- 交易支付中
	TS02 TransStatus = "02" //"02"-- 交易成功
	TS03 TransStatus = "03" //"03"-- 交易失败
	TS04 TransStatus = "04" //"04"-- 交易关闭
	TS05 TransStatus = "05" //"05"-- 已撤销
	TS10 TransStatus = "10" //"10"-- 已退款
	TS11 TransStatus = "11" //"11"-- 退款中
	TS12 TransStatus = "12" //"12"-- 退款异常
	TS13 TransStatus = "13" //"13"-- 退款关闭
)
