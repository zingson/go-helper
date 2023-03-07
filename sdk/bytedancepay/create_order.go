package bytedancepay

import (
	"encoding/json"
	"errors"
	"fmt"
)

//	服务端预下单  文档：https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/ecpay/server-doc
func CreateOrder(conf *Config, params *CreateOrderParams) (data *CreateOrderResultData, err error) {
	pmap := structToMap(params)
	pmap["sign"] = getSign(pmap, conf.Salt)
	rBytes, err := json.Marshal(pmap)
	if err != nil {
		return
	}
	resBody, err := post(conf, "/api/apps/ecpay/v1/create_order", string(rBytes))
	if err != nil {
		return
	}
	var rs *CreateOrderResult
	err = json.Unmarshal([]byte(resBody), &rs)
	if err != nil {
		return
	}
	if rs.ErrNo != 0 {
		err = errors.New(fmt.Sprintf("Z%d:%s", rs.ErrNo, rs.ErrTips))
		return
	}
	data = rs.Data
	return
}

type CreateOrderParams struct {
	AppId        string `json:"app_id"`        //必填，小程序appid
	OutOrderNo   string `json:"out_order_no"`  //必填，开发者侧的订单号, 同一小程序下不可重复
	TotalAmount  int64  `json:"total_amount"`  //必填，价格; 接口中参数支付金额单位为[分]
	Subject      string `json:"subject"`       //必填，商品描述; 长度限制 128 字节，不超过 42 个汉字
	Body         string `json:"body"`          //必填，商品详情
	ValidTime    int64  `json:"valid_time"`    //必填，订单过期时间(秒); 最小 15 分钟，最大两天
	CpExtra      string `json:"cp_extra"`      //非必填，开发者自定义字段，回调原样回传
	NotifyUrl    string `json:"notify_url"`    //非必填，商户自定义回调地址
	ThirdpartyId string `json:"thirdparty_id"` //非必填，第三方平台服务商 id，非服务商模式留空
	DisableMsg   int64  `json:"disable_msg"`   //非必填，是否屏蔽担保支付的推送消息，1-屏蔽 0-非屏蔽，接入 POI 必传
	MsgPage      string `json:"msg_page"`      //非必填，担保支付消息跳转页
}

type CreateOrderResult struct {
	ErrNo   int64                  `json:"err_no"`
	ErrTips string                 `json:"err_tips"`
	Data    *CreateOrderResultData `json:"data"`
}
type CreateOrderResultData struct {
	OrderId    string `json:"order_id"`    // 字节订单编号
	OrderToken string `json:"order_token"` // 用于小程序调用支付接口
}

func (m *CreateOrderResultData) Json() string {
	b, _ := json.Marshal(&m)
	return string(b)
}
