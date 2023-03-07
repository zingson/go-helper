package test

import (
	"root/src/sdk/upapi"
	"testing"
	"time"
)

/**
 * 5.70
 * 会员中心查询操作(memberPointBalance)
 * 1. 测试信息出现异常
	请求URL：https://open.95516.com/open/access/1.0/memberPointBalance
	请求报文：{"appId":"2d0851ef561945f4beaf68709601e6ef","backendToken":"wRVSS1t4SqmQPBJP5nbIFg==","mobile":"96oJNwI/JTQTP9ypoVspeg==","openId":"","sysId":"1026","transSeqId":"035025460d4f04b749655e0ba4495624","transTs":"20220309"}
	响应报文：{"resp":"02","msg":"column error acctEntityId","params":{}}

*/
func TestMemberPointBalance(t *testing.T) {
	result, err := upapi.MemberPointBalance(cfgtoml(), &upapi.MemberPointBalanceParams{
		SysId:        "1026",
		Mobile:       "13611703040",
		TransSeqId:   upapi.Rand32(),
		TransTs:      time.Now().Format("20060102"),
		BackendToken: "wRVSS1t4SqmQPBJP5nbIFg==",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("\n\n会员中心查询操作:", result, "\n\n")
}
