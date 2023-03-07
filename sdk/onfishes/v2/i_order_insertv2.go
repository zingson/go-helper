package v2

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// 3.1订单提交接口   http://*:*/Order/InsertV2（从供货商处获取）
func (c *Client) OrderInsertV2(p *OrderInsertV2Params) (result *OrderInsertV2Result, err error) {
	timestamp := Timestamp()
	CustomerIP := "127.0.0.1"
	extendParam, _ := json.Marshal(p.ExtendParam)
	sign := Md5Sign(c.Cfg.AppKey + strconv.FormatInt(p.BuyCount, 10) + p.CallBackUrl + p.ChargeAccount + CustomerIP + p.MOrderID + p.MemberAmountCode + p.ProductCode + timestamp + c.Cfg.AppSecret)
	err = c.Request("/Order/InsertV2",
		fmt.Sprintf("CallBackUrl=%s&ChargeAccount=%s&CustomerIP=%s&MemberAmountCode=%s&AppKey=%s&Sign=%s&MOrderID=%s&TimesTamp=%s&BuyCount=%s&ProductCode=%s&ExtendParam=%s&Attach=",
			p.CallBackUrl, p.ChargeAccount, CustomerIP, p.MemberAmountCode, c.Cfg.AppKey, sign, p.MOrderID, timestamp, strconv.FormatInt(p.BuyCount, 10), p.ProductCode, extendParam), &result)
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

type OrderInsertV2Params struct {
	MemberAmountCode string
	ProductCode      string                          //购买产品的编号（从服务商处获取）
	BuyCount         int64                           //需要购买的数量
	MOrderID         string                          // 商户平台订单号（唯一标识）
	ChargeAccount    string                          //充值到帐的账号
	ExtendParam      *OrderInsertV2ParamsExtendParam // 不同业务对应不同的扩展参数（详见：四、扩展参数说明）
	CallBackUrl      string                          //服务商会将充值结果推送到订单提交对应的地址，如果不传将不进行结果推送（推送方式详见：3.2订单回调功能）
}

type OrderInsertV2ParamsExtendParam struct {
	ExtendAccount string `json:"extendAccount,omitempty"` // 扩展账号用于传入除充值账号account外的其他账号，例如加油卡充值时的持卡人手机号码
}

type OrderInsertV2Result struct {
	Code        int64
	Msg         string
	TimesTamp   int64
	Sign        string
	OrderID     string
	ExtendParam *OrderInsertV2ResultExtendParam
}

type OrderInsertV2ResultExtendParam struct {
}
