package test

import (
	"root/src/sdk/upapi"
	"testing"
	"time"
)

/**
 * 5.72
 * 会员中心积点核销操作(memberPointDeduct)
 * 1. 测试出现异常
	请求URL：https://open.95516.com/open/access/1.0/memberPointDeduct
	请求报文：{"appId":"2d0851ef561945f4beaf68709601e6ef","costSource":"","descCode":"tt","mobile":"","nonceStr":"WpT32EcwPt","openId":"AwwA3GYyDmuLcX1M/1AKbi2WngbhguDCqeDiHaAWQR0qa/JmNYBXjhPWqPa+Rdbd","pointAt":"37wc8euJzFc=","signature":"PCwUlxKMptAFQD5UEyMQ17tAmFv55aItyu0HiRHQoN4O/cyhragNdfcwdwcRO1ig7a8wBRN/r+4u1Z8DDUpVjfrnG+K+ca0GBeWhYs7je8Cz2+wq31CVw1vdx66hJPCQqTATTEkdJQDLACL8uBiFuSabxZTUFVGxqVx8Vo7xrRQ=","sysId":"1026","timestamp":"1646814784","transDigest":"积点核销测试","transSeqId":"dc6dd14ec6f65c6a7b5f7a3c5d6fbbcf","transTs":"20220309043304"}
	响应报文：{"resp":"BC0025","msg":"签名验证失败","params":{}}
*/
func TestMemberPointDeduct(t *testing.T) {
	result, err := upapi.MemberPointDeduct(cfgtoml(), &upapi.MemberPointDeductParams{
		SysId:  "1026",
		OpenId: "AwwA3GYyDmuLcX1M/1AKbi2WngbhguDCqeDiHaAWQR0qa/JmNYBXjhPWqPa+Rdbd",
		//Mobile:      "13611703040",
		TransSeqId:  upapi.Rand32(),
		TransTs:     time.Now().Format("20060102030405"),
		PointAt:     "1",
		TransDigest: "积点核销测试",
		DescCode:    "tt",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("\n\n会员中心积点核销操作:", result, "\n\n")
}
