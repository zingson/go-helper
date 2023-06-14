package bank_zsyh

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

// RefundAPI 退款 接口文档:http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/RefundAPI
// 生产环境 https://merchserv.netpay.cmbchina.com/merchserv/BaseHttp.dll?DoRefundV2
// 测试环境 http://merchserv.cmburl.cn/merchserv/BaseHttp.dll?DoRefundV2
func RefundAPI(conf *Config, req *RefundReq) (res *RefundRes, err error) {
	//reqMap := StructToMap(req.ReqData)
	//waitForSignStr := SortMap(reqMap, true) + "&" + conf.Merkey
	//req.Sign = Sha256Sign(waitForSignStr)

	rBytes, err := json.Marshal(req)
	if err != nil {
		return
	}
	resBody, err := PostForm(conf, conf.MerchservUrl, string(rBytes))
	if err != nil {
		return
	}
	logrus.Infof("QueryFbPubKey, response:%s", resBody)
	res = &RefundRes{}
	err = json.Unmarshal([]byte(resBody), res)
	if err != nil {
		return
	}
	if res.RspData.RspCode != "SUC0000" {
		err = errors.New(fmt.Sprintf("rspCode:%s, rspMsg:%s", res.RspData.RspCode, res.RspData.RspMsg))
		return
	}
	return res, nil
}

type RefundReq struct {
	ReqData *RefundReqData `json:"reqData"`
}

type RefundReqData struct {
	DateTime       string `json:"dateTime"`       //商户发起该请求的当前时间，精确到秒。 格式：yyyyMMddHHmmss
	BranchNo       string `json:"branchNo"`       //商户分行号，4位数字
	MerchantNo     string `json:"merchantNo"`     //商户号，6位数字
	Date           string `json:"date"`           //商户订单日期，支付时的订单日期 格式：yyyyMMdd
	OrderNo        string `json:"orderNo"`        //商户订单号，支付时的订单号
	RefundSerialNo string `json:"refundSerialNo"` //退款流水号,商户生成，同一笔订单内，同一退款流水号只能退款一次。可用于防重复退款。
	Amount         string `json:"amount"`         //退款金额,格式xxxx.xx
	Desc           string `json:"desc"`           //退款描述
}

type RefundRes struct {
	RspData struct {
		RspCode        string `json:"rspCode"`        //处理结果,SUC0000：请求处理成功 其他：请求处理失败
		RspMsg         string `json:"rspMsg"`         //详细信息,请求处理失败时返回错误描述
		DateTime       string `json:"dateTime"`       //响应时间,银行返回该数据的时间 格式：yyyyMMddHHmmss
		BankSerialNo   string `json:"bankSerialNo"`   //银行的退款流水号
		Currency       string `json:"currency"`       //退款币种,固定为：“10”
		Amount         string `json:"amount"`         //退款金额,格式：xxxx.xx
		RefundRefNo    string `json:"refundRefNo"`    //银行的退款参考号
		BankDate       string `json:"bankDate"`       //退款受理日期 格式：yyyyMMdd
		BankTime       string `json:"bankTime"`       //退款受理时间 格式：HHmmss
		RefundSerialNo string `json:"refundSerialNo"` //商户上送流水号
		SettleAmount   string `json:"settleAmount"`   //实际退款金额,格式：xxxx.xx
		DiscountAmount string `json:"discountAmount"` //退回优惠金额,格式：xxxx.xx
	} `json:"rspData"`
}
