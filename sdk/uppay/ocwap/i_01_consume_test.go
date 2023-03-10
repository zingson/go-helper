package ocwap

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

// 消费测试
func TestMain1(t *testing.T) {
	// https://msd.himkt.cn/work/consume?orderId=T0000001&txnAmt=1&accNo=6251211100976741
	// https://msd.himkt.cn/work/consume?orderId=T0000002&txnAmt=1&accNo=6214830213065526
	http.HandleFunc("/consume", func(writer http.ResponseWriter, request *http.Request) {
		// 跳转银联全渠道手机网页支付界面
		err := Consume(_cfg(), &ConsumeParams{
			AccNo:       request.FormValue("accNo"),
			OrderId:     request.FormValue("orderId"),
			TxnAmt:      request.FormValue("txnAmt"),
			FrontUrl:    "https://msd.himkt.cn/work/consume/front",
			BackUrl:     "https://msd.himkt.cn/work/consume/back",
			TxnTime:     TxnTime(),
			ReqReserved: "--",
		}, writer)
		if err != nil {
			return
		}
	})

	// 前端接受通知，判断状态00，跳转成功节目
	http.Handle("/consume/front", ConsumeNotifyFrontHandler(func(writer http.ResponseWriter, request *http.Request, entity *ConsumeNotifyEntity, err error) {
		if err != nil {
			fmt.Println(err.Error())
			writer.Write([]byte(err.Error()))
			return
		}
		v, _ := json.Marshal(entity)

		writer.Write([]byte(request.RequestURI + " <br>"))
		writer.Write(v)
	}))

	// 红土哎接受通知，判断状态00，调用查询接口，曲儿oriRespCode等于00，执行业务发货逻辑
	http.Handle("/consume/back", ConsumeNotifyHandler(func(o *ConsumeNotifyEntity) error {
		bytes, _ := json.Marshal(o)
		fmt.Println("JSON ConsumeNotifyEntity ", string(bytes))
		return nil
	}))

	// https://msd.himkt.cn/work/query?orderId=T0000002
	http.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		orderId := request.FormValue("orderId")
		result, err := Query(_cfg(), orderId)
		if err != nil {
			fmt.Println("Query Error", err.Error())
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Println("resultJSON", string(bytes))
		writer.Write(bytes)
	})

	fmt.Println("Start serve Listen 80 ...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
