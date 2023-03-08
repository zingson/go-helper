package cmbnetpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/htime"
)

// QueryFbPubKey 查询招行公钥 接口文档:http://openhome.cmbchina.com/PayNew/pay/doc/cell/H5/QueryKeyAPI
// 生产环境 https://b2b.cmbchina.com/CmbBank_B2B/UI/NetPay/DoBusiness.ashx
// 测试环境 https://cmbchinab2b.bas.cmburl.cn/CmbBank_B2B/UI/NetPay/DoBusiness.ashx
func QueryFbPubKey(conf *Config) (fbPubKey string, err error) {
	req := &QueryFbPubKeyReq{FixedParams: GetFixedParams()}
	req.ReqData.DateTime = htime.NowF14()
	req.ReqData.TxCode = "FBPK"
	req.ReqData.BranchNo = conf.BranchNo
	req.ReqData.MerchantNo = conf.MerchantNo

	reqMap := StructToMap(req.ReqData)
	waitForSignStr := SortMap(reqMap, true) + "&" + conf.Merkey
	req.Sign = Sha256Sign(waitForSignStr)

	rBytes, err := json.Marshal(req)
	if err != nil {
		return
	}
	resBody, err := PostForm(conf, conf.ApiUrl, string(rBytes))
	if err != nil {
		return
	}
	logrus.Infof("QueryFbPubKey, response:%s", resBody)
	res := &QueryFbPubKeyRes{}
	err = json.Unmarshal([]byte(resBody), res)
	if err != nil {
		return
	}
	if res.RspData.RspCode != "SUC0000" {
		err = errors.New(fmt.Sprintf("rspCode:%s, rspMsg:%s", res.RspData.RspCode, res.RspData.RspMsg))
		return
	}
	return res.RspData.FbPubKey, nil
}

type QueryFbPubKeyReq struct {
	FixedParams
	ReqData struct {
		DateTime   string `json:"dateTime"`   //商户发起该请求的当前时间，精确到秒 格式：yyyyMMddHHmmss
		TxCode     string `json:"txCode"`     //交易码,固定为“FBPK”
		BranchNo   string `json:"branchNo"`   //商户分行号，4位数字
		MerchantNo string `json:"merchantNo"` //商户号，6位数字
	} `json:"reqData"`
}

type QueryFbPubKeyRes struct {
	FixedParams
	RspData struct {
		RspCode  string `json:"rspCode"`  //处理结果,SUC0000：请求处理成功 其他：请求处理失败
		RspMsg   string `json:"rspMsg"`   //详细信息,失败时返回具体失败原因
		FbPubKey string `json:"fbPubKey"` //用Base64编码的招行公钥
		DateTime string `json:"dateTime"` //响应时间,银行返回该数据的时间 格式：yyyyMMddHHmmss
	} `json:"rspData"`
}
