package bank_zhywt

import (
	"errors"
	"fmt"
	"time"
)

// QueryKeyAPI 查询招行公钥API 文档:http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/QueryKeyAPI
func QueryKeyAPI(conf *Config) (fbPubKey string, err error) {
	rspData, err := Post[QueryKeyAPIReq, QueryKeyAPIRes](conf, conf.CmbBankB2BUrl+"/CmbBank_B2B/UI/NetPay/DoBusiness.ashx", QueryKeyAPIReq{
		DateTime:   time.Now().Local().Format("20060102150405"),
		TxCode:     "FBPK",
		BranchNo:   conf.BranchNo,
		MerchantNo: conf.MerchantNo,
	})
	if err != nil {
		return
	}
	if rspData.RspCode != SUC0000 {
		err = errors.New(fmt.Sprintf("%s:%s", rspData.RspCode, rspData.RspMsg))
		return
	}
	fbPubKey = rspData.FbPubKey
	return
}

type QueryKeyAPIReq struct {
	DateTime   string `json:"dateTime"`   //商户发起该请求的当前时间，精确到秒 格式：yyyyMMddHHmmss
	TxCode     string `json:"txCode"`     //交易码,固定为“FBPK”
	BranchNo   string `json:"branchNo"`   //商户分行号，4位数字
	MerchantNo string `json:"merchantNo"` //商户号，6位数字
}

type QueryKeyAPIRes struct {
	RspCode  string `json:"rspCode"`  //处理结果,SUC0000：请求处理成功 其他：请求处理失败
	RspMsg   string `json:"rspMsg"`   //详细信息,失败时返回具体失败原因
	FbPubKey string `json:"fbPubKey"` //用Base64编码的招行公钥
	DateTime string `json:"dateTime"` //响应时间,银行返回该数据的时间 格式：yyyyMMddHHmmss
}
