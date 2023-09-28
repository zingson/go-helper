package test

import (
	"github.com/zingson/go-helper/sdk/upapi"
	"testing"
	"time"
)

/**
 * 5.71
 * 会员中心积点赠送操作(memberPointAcquire)
 *
 * 测试openId
 * 示例1:AwwA3GYyDmuLcX1M/1AKbi2WngbhguDCqeDiHaAWQR0qa/JmNYBXjhPWqPa+Rdbd
 * 示例2:N8qOwIpwrvhZTQ81S54TDKAA/o/vIEQJQ826Pp2UYh//EfrLsRNOMCSFgejMd79k
 */
func TestMemberPointAcquire(t *testing.T) {
	result, err := upapi.MemberPointAcquire(cfgtoml(), &upapi.MemberPointAcquireParams{
		SysId: "6010",
		//OpenId: "AwwA3GYyDmuLcX1M/1AKbi2WngbhguDCqeDiHaAWQR0qa/JmNYBXjhPWqPa+Rdbd",
		Mobile:      "13611703040",
		GetSource:   "827072017227587584",
		TransSeqId:  upapi.Rand32(),
		TransTs:     time.Now().Format("20060102030405"),
		PointAt:     "5",
		TransDigest: "甬城天天惠",
		DescCode:    "827072017227587584",
		InMode:      "0",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("\n\n会员中心积点赠送操作:", result, "\n\n")
}
