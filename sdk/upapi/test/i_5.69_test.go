package test

import (
	"root/src/sdk/upapi"
	"testing"
	"time"
)

/**
 * 5.69
 * 给用户赠送62VIP会员操作(memberVipAcquire)
 * 1. 注意，测试信息出现异常
 	请求URL：https://open.95516.com/open/access/1.0/memberVipAcquire
	请求报文：{"appId":"2d0851ef561945f4beaf68709601e6ef","isLimit":"0","memberType":"01","mobile":"96oJNwI/JTQTP9ypoVspeg==","nonceStr":"BcphWiS8l9","openId":"","signature":"oP0fK6kMoVPxX2mqHDeo39yFQzKGNqn6xKyk7/BpNK9yahvhPGd5HIlurODj6/9ntlhfEFZat35KHYesrkIBBF66k2fQHiZADsZFK5TjajnWCiv48zwx5iVlr3YpIDMYhMgPPswvXRH+aGgrsAoTSOJb19NX3aQ0vbnbo9pfHS0=","sysId":"1026","timestamp":"1646814278","transSeqId":"d3c79759139fb5fc7f61c52f432c02b1","transTs":"20220309042438"}
	响应报文：{"resp":"02","msg":"column error acctEntityId","params":{}}
*/
func TestMemberVipAcquire(t *testing.T) {
	result, err := upapi.MemberVipAcquire(cfgtoml(), &upapi.MemberVipAcquireParams{
		SysId:      "1026",
		Mobile:     "13611703040",
		MemberType: "01",
		TransSeqId: upapi.Rand32(),
		TransTs:    time.Now().Format("20060102030405"),
		IsLimit:    "0",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("\n\n给用户赠送62VIP会员操作:", result, "\n\n")
}
