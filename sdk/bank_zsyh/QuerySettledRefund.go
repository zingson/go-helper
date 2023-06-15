package bank_zsyh

import (
	"errors"
	"fmt"
)

// QuerySettledRefund 文档：http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/QuerySettledRefund
func QuerySettledRefund(conf *Config, reqData QuerySettledRefundReq) (rspData QuerySettledRefundRes, err error) {
	rspData, err = Post[QuerySettledRefundReq, QuerySettledRefundRes](conf, conf.MerchservUrl+"/merchserv/BaseHttp.dll?QuerySettledRefundV2", reqData)
	if err != nil {
		return
	}
	if rspData.RspCode != SUC0000 {
		err = errors.New(fmt.Sprintf("%s:%s", rspData.RspCode, rspData.RspMsg))
		return
	}
	return
}

type QuerySettledRefundReq struct {
	DateTime         string `json:"dateTime"`         //
	BranchNo         string `json:"branchNo"`         //
	MerchantNo       string `json:"merchantNo"`       //
	Type             string `json:"type"`             //查询类型 A：按银行退款流水号查单笔 需上送：bankSerialNo,date   B：按商户订单号+商户退款流水号+订单日期查单笔 需上送：orderNo,date,merchantSerialNo  C: 按商户订单号查对应的所有退款 需上送以：orderNo,,date
	BankSerialNo     string `json:"bankSerialNo"`     //银行退款流水号长度不超过20位
	Date             string `json:"date"`             //商户原支付交易的订单日期，格式：yyyyMMdd
	OrderNo          string `json:"orderNo"`          //商户原支付交易的订单号
	MerchantSerialNo string `json:"merchantSerialNo"` //商户退款流水号长度不超过20位
}

type QuerySettledRefundRes struct {
	RspCode   string `json:"rspCode"`   //处理结果,SUC0000：请求处理成功 其他：请求处理失败
	RspMsg    string `json:"rspMsg"`    //详细信息,请求处理失败时返回错误描述
	DateTime  string `json:"dateTime"`  //响应时间,银行返回该数据的时间 格式：yyyyMMddHHmmss
	DataCount string `json:"dataCount"` //返回条数
	DataList  string `json:"dataList"`  // 每笔记录一行，每行以\r\n 结束。第一行为表头;从第二行起为数据记录;各参数以逗号和`符号分隔(`为标准键盘1 左边键的字符)，字段顺序与表头一致。 表单字段说明查看接口文档
}
