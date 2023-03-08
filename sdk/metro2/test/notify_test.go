package test

import (
	"context"
	"github.com/zingson/go-helper/h"
	"github.com/zingson/go-helper/sdk/metro2"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	k := metro2.TokenKey("testtoken")
	t.Log(k)
	h.REDIS.Set(context.TODO(), k, "{\"GuidUser\":\"13611703040\",\"MobilePhone\":\"13611703040\"}", 12*time.Hour)
}

func TestQueryUser(t *testing.T) {
	//https://msd.himkt.cn/gw/voucher/metro2/user/info
	config.ServiceUrl = "http://localhost"
	resData, err := metro2.HttpPost(config, metro2.Rand32(), "/metro2/user/info", "{\"AppToken\":\"testtoken\"}")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resData)
}
