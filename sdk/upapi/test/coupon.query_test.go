package test

import (
	"fmt"
	"root/src/sdk/upapi"
	"testing"
)

// 5.8.10  赠送优惠券结果查询 <coupon.query>
func Test_CouponQuery(t *testing.T) {
	//"transSeqId":"4ff007f90a384eee869d99d7166ed342","transTs":"20200906"
	err := upapi.CouponQuery(cfgtoml(), upapi.Rand32(), "2922b52272a04cc9a6cec290c2a6d324", "20201015", func(config *upapi.Config) string {
		return ""
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestFmt(t *testing.T) {
	t.Log(fmt.Sprintf("xxxx : %s", "123"))
	t.Log(fmt.Sprintf("xxxx : %v", "123"))
}
