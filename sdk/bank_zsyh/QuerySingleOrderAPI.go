package bank_zsyh

import (
	"errors"
	"fmt"
)

// QuerySingleOrderAPI 查询单笔订单 接口文档:http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/QuerySingleOrderAPI
func QuerySingleOrderAPI(conf *Config, reqData QuerySingleOrderAPIReq) (rspData QuerySingleOrderAPIRes, err error) {
	rspData, err = Post[QuerySingleOrderAPIReq, QuerySingleOrderAPIRes](conf, conf.MerchservUrl+"/merchserv/BaseHttp.dll?QuerySingleOrder", reqData)
	if err != nil {
		return
	}
	if rspData.RspCode != SUC0000 {
		err = errors.New(fmt.Sprintf("%s:%s", rspData.RspCode, rspData.RspMsg))
		return
	}
	return
}

type QuerySingleOrderAPIReq struct {
	DateTime     string `json:"dateTime"`     //商户发起该请求的当前时间，精确到秒。 格式：yyyyMMddHHmmss
	BranchNo     string `json:"branchNo"`     //商户分行号，4位数字
	MerchantNo   string `json:"merchantNo"`   //商户号，6位数字
	Date         string `json:"date"`         //商户订单日期，支付时的订单日期 格式：yyyyMMdd
	OrderNo      string `json:"orderNo"`      //商户订单号，支付时的订单号
	Type         string `json:"type"`         //查询类型，A：按银行订单流水号查询 B：按商户订单日期和订单号查询；
	BankSerialNo string `json:"bankSerialNo"` //银行订单流水号,type=A时必填
}
type QuerySingleOrderAPIRes struct {
	RspCode            string      `json:"rspCode"`            //处理结果,SUC0000：请求处理成功 其他：请求处理失败
	RspMsg             string      `json:"rspMsg"`             //详细信息,请求处理失败时返回错误描述
	DateTime           string      `json:"dateTime"`           //响应时间,银行返回该数据的时间 格式：yyyyMMddHHmmss
	BankSerialNo       string      `json:"bankSerialNo"`       //银行的退款流水号
	Currency           string      `json:"currency"`           //退款币种,固定为：“10”
	Amount             string      `json:"amount"`             //退款金额,格式：xxxx.xx
	QueryOrderRefNo    string      `json:"QueryOrderRefNo"`    //银行的退款参考号
	BankDate           string      `json:"bankDate"`           //退款受理日期 格式：yyyyMMdd
	BankTime           string      `json:"bankTime"`           //退款受理时间 格式：HHmmss
	QueryOrderSerialNo string      `json:"QueryOrderSerialNo"` //商户上送流水号
	SettleAmount       string      `json:"settleAmount"`       //实际退款金额,格式：xxxx.xx
	DiscountAmount     string      `json:"discountAmount"`     //退回优惠金额,格式：xxxx.xx
	BranchNo           string      `json:"branchNo"`           //商户分行号，4位数字
	MerchantNo         string      `json:"merchantNo"`         //商户号，6位数字
	Date               string      `json:"date"`               //商户订单日期，格式：yyyyMMdd
	OrderNo            string      `json:"orderNo"`            //商户订单号
	OrderAmount        string      `json:"orderAmount"`        //订单金额,格式：xxxx.xx
	Fee                string      `json:"fee"`                //费用金额,格式：xxxx.xx
	OrderStatus        OrderStatus `json:"orderStatus"`        //订单状态, 0:已结帐 1:已撤销 2:部分结账 4:未结帐 6:未知状态/订单失败 7:冻结交易—冻结金额已经全部结账 8:冻结交易，冻结金额只结账了一部分  0代表终态
	SettleDate         string      `json:"settleDate"`         //银行处理日期,格式：yyyyMMdd
	SettleTime         string      `json:"settleTime"`         //银行处理时间,格式：HHmmss
	CardType           string      `json:"cardType"`           //卡类型,02：本行借记卡； 03：本行贷记卡； 08：他行借记卡； 09：他行贷记卡
	MerchantPara       string      `json:"merchantPara"`       //商户自定义参数,商户在支付接口中传送的merchantPara参数，超过100字节自动截断。
}

// OrderStatus 订单状态
type OrderStatus string

const (
	OrderStatus0 OrderStatus = "0" //已结帐
	OrderStatus1 OrderStatus = "1" //已撤销
	OrderStatus2 OrderStatus = "2" //部分结账
	OrderStatus4 OrderStatus = "4" //未结帐
	OrderStatus6 OrderStatus = "6" //未知状态/订单失败
	OrderStatus7 OrderStatus = "7" //冻结交易—冻结金额已经全部结账
	OrderStatus8 OrderStatus = "8" //冻结交易，冻结金额只结账了一部分
)
