package example

import (
	"encoding/json"
	"github.com/zingson/go-helper/sdk/uppay/ocapp"
	"testing"
)

// 消费下单测试，获取tn，使用云闪付upsdk调起支付
func TestConsume(t *testing.T) {
	result, err := ocapp.Consume(cfg89833027372F284, &ocapp.ConsumeParams{
		OrderId:     "17003880709960300587100268223",
		TxnAmt:      1,
		BackUrl:     "https://msd.himkt.cn/gw/62vip/order/call/back",
		TxnTime:     ocapp.TxnTime(),
		ReqReserved: "-",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	rbytes, _ := json.Marshal(result)
	t.Log(string(rbytes))
}

func TestPost1(t *testing.T) {
	s := "bizType=000201&backUrl=https%3A%2F%2Fmsd.himkt.cn%2Fgw%2F62vip%2Forder%2Fcall%2Fback&orderId=17003880709960300587100268223&txnSubType=01&signature=G3iW40frQ542VF0n9wNUk%2BY8Al6VWzRwOGGnS8WMUaguGBGAFJve5NWtQlZeRxeDLT%2ByPssuYII3QZyF5Z5BHJT%2FtAYTwUfKb8Qm8eHDF5cw2Lb7%2Bqb4RfcqzmrF%2F6LyLIRPa%2B4%2F5ocbH2lPfV6E%2Bd6OBnoKsqxjy8%2FeWIy2%2BrQMwigCYX7gsXFy1T7UAcMw2Le44oLOAE69Kc34M%2B3QGvdIVS%2BwDakVdWg%2Bk6dWSE8MAI6hPWai8FtpAii0fWGg4KMfTMhAU0QO516%2BtSFdnkuhcpYSa9wnRscUAAOLymzlja7RVGF8I94l2%2FapYI%2BtVM6nhI2IyPnML9aY6ZorjA%3D%3D&channelType=08&txnType=01&certId=86842351990&encoding=UTF-8&version=5.1.0&accessType=0&reqReserved=-&txnTime=20231121153644&merId=89833027372F284&currencyCode=156&signMethod=01&txnAmt=1"
	ocapp.Post("consume", "https://gateway.95516.com/gateway/api/appTransReq.do", s)
}

/*

 */
