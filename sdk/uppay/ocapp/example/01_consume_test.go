package example

import (
	"encoding/json"
	"github.com/zingson/go-helper/sdk/uppay/ocapp"
	"testing"
)

// 消费下单测试，获取tn，使用云闪付upsdk调起支付
func TestConsume(t *testing.T) {
	result, err := ocapp.Consume(cfg89833027372F284, &ocapp.ConsumeParams{
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

func TestPost1(t *testing.T) {
	s := "bizType=000201&backUrl=https%3A%2F%2Fmsd.himkt.cn%2Fgw%2F62vip%2Forder%2Fcall%2Fback&orderId=17003880709960300587100268223&txnSubType=01&signature=eFUL%2Bw2bUmGQu1gUdVan%2BissqUe%2BXnWZUhhbzmK3WKNlSpzpEqSCdFni%2BfszQasmZxaneKvfVSBPGTmlxql3XE4tetnnjwGffH82SjsloJicSzOz9K1oa%2BvEOQZx0%2ByTE6wstu%2FqJutLxgPulGUNCmxEc0V3JEsUlBGpJQOXDBQnmIMeB7ga8897XSFvoIkMsiqiKJ%2Btbw6n4sY94hqijRRIBXmEM5cChv%2Fcx7QkEJ3YPIH95LgkgmDDf3wv0NjI7TIl8qRQuIa5wGIT9apyaRtflx6dGLnCTMJPAXe0WhJ9GM2M4Jx6nCDdEAyqc88HYISxT6o%2B%2BBVMlhJXkt0QKA%3D%3D&channelType=08&txnType=01&certId=81628889475&encoding=UTF-8&version=5.1.0&accessType=0&reqReserved=-&txnTime=20231121150414&merId=89833027372F284&currencyCode=156&signMethod=01&txnAmt=1"
	ocapp.Post("consume", "https://gateway.95516.com/gateway/api/appTransReq.do", s)
}

/*

 */
