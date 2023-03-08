package test

import (
	"github.com/zingson/go-helper/sdk/upapi"
	"testing"
)

// 机构账户（红包）余额查询
func TestRedPacketSelect(t *testing.T) {
	rs, err := upapi.RedPacketSelect(cfgtoml(), "P220427102525711", func(config *upapi.Config) string {
		r, err := upapi.BackendToken(config)
		if err != nil {
			panic(err)
		}
		return r.BackendToken
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rs.Json())
}
