package test

import (
	"root/src/sdk/upapi"
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
		SysId:  "1026",
		OpenId: "AwwA3GYyDmuLcX1M/1AKbi2WngbhguDCqeDiHaAWQR0qa/JmNYBXjhPWqPa+Rdbd",
		//Mobile:      "13611703040",
		GetSource:   "tt",
		TransSeqId:  upapi.Rand32(),
		TransTs:     time.Now().Format("20060102030405"),
		PointAt:     "1",
		TransDigest: "积点赠送测试2",
		DescCode:    "tt",
		InMode:      "0",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("\n\n会员中心积点赠送操作:", result, "\n\n")
}
