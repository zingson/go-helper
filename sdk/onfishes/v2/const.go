package v2

/*
1	订单正在处理中	等待并且继续查询（不能做成功和失败处理）
2	订单成功	做成功处理
3	订单失败	做失败处理
*/
type OrderStatus string

const (
	ORDER_STATUS_1 OrderStatus = "1"
	ORDER_STATUS_2 OrderStatus = "2"
	ORDER_STATUS_3 OrderStatus = "3"
)
