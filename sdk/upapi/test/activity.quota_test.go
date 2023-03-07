package test

import (
	"root/src/sdk/upapi"
	"testing"
)

// 5.8.12  优惠券活动剩余名额查询
func TestActivityQuota(t *testing.T) {
	_, err := upapi.ActivityQuota(cfgtoml(), upapi.Rand32(), "3102022100115856", "3", func(config *upapi.Config) string {
		bt, err := upapi.BackendToken(config)
		if err != nil {
			t.Log(err)
			return ""
		}
		return bt.BackendToken
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	// 测试结果：404 Not Found
}
