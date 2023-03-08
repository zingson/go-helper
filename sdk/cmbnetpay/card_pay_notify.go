package cmbnetpay

import (
	"fmt"
	"net/http"
)

// NotifyRSAVerify 成功支付结果验签
func NotifyRSAVerify(payNotify *PayNotify) (ok bool, err error) {
	//order, err := hd.NewOrderDao().FindOne(nil, bson.D{bson.E{"order_id", payNotify.NoticeData.OrderNo}})
	//if err != nil {
	//	err = errors.New(fmt.Sprintf("cmbnetpay:NotifyRSAVerify:查询订单失败, payNotify:%v, error:%s", payNotify, err.Error()))
	//	return
	//}
	//通过渠道编号查询渠道信息
	//chap, err := plat.NewChapDao().FindOne(nil, bson.D{bson.E{"chid", order.PayChid}})
	//if err != nil {
	//return
	//}

	//conf := &Config{}
	//err = utils.ConvertStructData(chap.Param[string(plat.CHAN_PAY_CMBNETPAY)], conf)
	//if err != nil {
	//	return
	//	}

	//pubkey, err := QueryFbPubKey(conf)
	//if err != nil {
	//	err = errors.New(fmt.Sprintf("cmbnetpay:NotifyRSAVerify:获取公钥失败, payNotify:%v, error:%s", payNotify, err.Error()))
	//	return
	//}
	//	reqMap := StructToMap(payNotify.NoticeData)
	//	waitForSignStr := SortMap(reqMap, true)
	//ok, err = RSAVerify(waitForSignStr, payNotify.Sign, []byte(pubkey))
	//if !ok {
	//	err = errors.New(fmt.Sprintf("cmbnetpay:NotifyRSAVerify:验签失败, payNotify:%v, error:%s", payNotify, err.Error()))
	//	return
	//	}
	return
}

// NotifyResponse 通知应答
func NotifyResponse(w http.ResponseWriter, err error) {
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code": "FAIL","message": "%s"}`, err.Error())))
		w.WriteHeader(500)
		return
	}
	w.Write([]byte(`{"code": "SUCCESS","message": "成功"}`))
	w.WriteHeader(200)
}

// PayNotify 成功支付结果
type PayNotify struct {
	FixedParams
	NoticeData struct {
		DateTime       string `json:"dateTime"`       //商户发起该请求的当前时间，精确到秒。 格式：yyyyMMddHHmmss
		Date           string `json:"date"`           //商户订单日期，支付时的订单日期 格式：yyyyMMdd
		Amount         string `json:"amount"`         //订单金额，格式：XXXX.XX
		BankDate       string `json:"bankDate"`       //银行受理日期
		OrderNo        string `json:"orderNo"`        //商户订单号
		DiscountAmount string `json:"discountAmount"` //单位为元，精确到小数点后两位。格式为：xxxx.xx元
		NoticeType     string `json:"noticeType"`     //本接口固定为：“BKPAYRTN”
		HttpMethod     string `json:"httpMethod"`     //固定为“POST”
		CardType       string `json:"cardType"`       //卡类型,02：本行借记卡； 03：本行贷记卡； 08：他行借记卡； 09：他行贷记卡
		NoticeSerialNo string `json:"noticeSerialNo"` //银行通知序号,订单日期+订单号
		MerchantPara   string `json:"merchantPara"`   //原样返回商户在一网通支付请求报文中传送的成功支付结果通知附加参数
		DiscountFlag   string `json:"discountFlag"`   //优惠标志,Y:有优惠 N：无优惠
		BankSerialNo   string `json:"bankSerialNo"`   //银行订单流水号
		NoticeUrl      string `json:"noticeUrl"`      //回调HTTP地址,支付请求时填写的支付结果通知地址
		BranchNo       string `json:"branchNo"`       //商户分行号，4位数字
		MerchantNo     string `json:"merchantNo"`     //商户号，6位数字
	} `json:"noticeData"`
}
