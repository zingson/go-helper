package main

import (
	"encoding/json"
	"github.com/zingson/go-helper/sdk/uppay/ocapp"
	"testing"
)

// 消费下单测试，获取tn，使用云闪付upsdk调起支付
func TestConsume(t *testing.T) {
	result, err := ocapp.Consume(cfg821330248164060, &ocapp.ConsumeParams{
		OrderId:     "30000000001112",
		TxnAmt:      1,
		BackUrl:     "https://msd.himkt.cn/work/consume/back",
		TxnTime:     ocapp.TxnTime(),
		ReqReserved: "",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	rbytes, _ := json.Marshal(result)
	t.Log(string(rbytes))
}
