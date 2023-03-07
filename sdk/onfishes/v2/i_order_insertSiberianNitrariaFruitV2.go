package v2

import (
	"fmt"
	"strconv"
)

// 3.5卡密申请接口  http://*:*/Order/InsertSiberianNitrariaFruitV2（请从供货商处获取）
func (c *Client) OrderInsertSiberianNitrariaFruitV2(p *OrderInsertSiberianNitrariaFruitV2Params) (result *OrderInsertV2Result, err error) {
	timestamp := Timestamp()
	CustomerIP := "127.0.0.1"
	extendParam := fmt.Sprintf("{\"MemberPublicKey\":\"%s\"}", c.Cfg.RsaPubKey)
	sign := Md5Sign(c.Cfg.AppKey + strconv.FormatInt(p.BuyCount, 10) + p.CallBackUrl + p.MOrderID + p.ProductCode + timestamp + c.Cfg.AppSecret)
	err = c.Request("/Order/InsertSiberianNitrariaFruitV2",
		fmt.Sprintf(
			"CallBackUrl=%s&ChargeAccount=%s&CustomerIP=%s&MemberAmountCode=%s&AppKey=%s&Sign=%s&MOrderID=%s&TimesTamp=%s&BuyCount=%s&ProductCode=%s&ExtendParam=%s&Attach=",
			p.CallBackUrl, p.ChargeAccount, CustomerIP, p.MemberAmountCode, c.Cfg.AppKey, sign, p.MOrderID, timestamp, strconv.FormatInt(p.BuyCount, 10), p.ProductCode, extendParam),
		&result)
	if err != nil {
		return
	}
	if result.Code != ERR_SUCCESS.Code {
		err = Err(result.Code, result.Msg)
		return
	}
	if result.Sign != Md5Sign(c.Cfg.AppKey+strconv.FormatInt(result.TimesTamp, 10)+strconv.FormatInt(result.Code, 10)+result.OrderID+c.Cfg.AppSecret) {
		err = ERR_SIGN
		return
	}
	return
}

type OrderInsertSiberianNitrariaFruitV2Params struct {
	MemberAmountCode string
	ProductCode      string //购买产品的编号（从服务商处获取）
	BuyCount         int64  //需要购买的数量 目前只支持1
	MOrderID         string // 商户平台订单号（唯一标识）
	ChargeAccount    string //充值到帐的账号 ，非必填
	CallBackUrl      string //服务商会将充值结果推送到订单提交对应的地址，如果不传将不进行结果推送（推送方式详见：3.2订单回调功能）
}

type OrderInsertSiberianNitrariaFruitV2Result struct {
	Code        int64
	Msg         string
	TimesTamp   int64
	Sign        string
	OrderID     string
	ExtendParam *OrderInsertSiberianNitrariaFruitV2ResultExtendParm
}

type OrderInsertSiberianNitrariaFruitV2ResultExtendParm struct {
	MOrderId     string
	ProductCode  string
	BuyCount     int64
	ProductPrice int64
	SellPrice    int64
	CallBackUrl  string
	TimesTamp    int64
	FinishTime   string
}
