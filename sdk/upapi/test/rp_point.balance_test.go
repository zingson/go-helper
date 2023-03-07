package test

import (
	"root/src/sdk/upapi"
	"testing"
)

//专享红包余额查询
//根据专享红包活动 id 查询用户的该专享红包余额。
func TestPointBalance(t *testing.T) {
	_, err := upapi.PointBalance(cfgtoml(), &upapi.PointBalanceParams{
		PointId:      "4122042628792805",
		AcctEntityTp: upapi.AETP_01,
		AcctEntityId: "13611703040",
		AccessId:     "",
	}, func(config *upapi.Config) string {
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
	return
}
