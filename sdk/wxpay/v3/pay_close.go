package v3

/*
以下情况需要调用关单接口：
1、商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
2、系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。

注意：
• 关单没有时间限制，建议在订单生成后间隔几分钟（最短5分钟）再调用关单接口，避免出现订单状态同步不及时导致关单失败。
*/

// 关闭订单API
func (pay *PayService) Close(mchid, out_trade_no string) (err error) {
	err = pay.HttpPost("/v3/pay/transactions/out-trade-no/"+out_trade_no+"/close", map[string]string{"mchid": mchid, "out_trade_no": out_trade_no}, nil)
	return
}
