package banknbcb

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// 支付通知处理
func PayNotify(request *http.Request, writer http.ResponseWriter, getConfig func(traceNo string) (cfg *Config, err error), f func(params *PayNotifyParams) error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(`{"return_code":"500","return_msg":"` + err.Error() + `"}`))
			logrus.Infof("hnbcb 支付通知处理异常 %s", err.Error())
			return
		}
		writer.WriteHeader(200)
		writer.Write([]byte(`{"return_code":"00","return_msg":"success"}`))
		logrus.Infof("hnbcb 支付通知处理成功")
	}()

	// 接收参数
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	if bytes == nil {
		err = errors.New("NBCB_WXPAY_NOTIFY_ERR:放心充 易收宝微信JSAPI支付通知参数为空")
		return
	}
	body := string(bytes)
	logrus.Infof("hnbcb 支付通知报文：%s", body)

	p := make(map[string]interface{})
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		return
	}
	traceNo := p["traceNo"].(string)

	cfg, err := getConfig(traceNo)
	if err != nil {
		return
	}

	// 验签
	sign := p["sign"].(string)
	p["sign"] = ""
	err = RsaVery(sortStr(p), sign, cfg.NbcbPubKey)
	if err != nil {
		return
	}

	// 解析
	params := &PayNotifyParams{}
	err = json.Unmarshal(bytes, &params)
	if err != nil {
		return
	}

	err = f(params)
	if err != nil {
		return
	}
	return
}

type PayNotifyParams struct {
	OutTradeId    string      `json:"outTradeId"`
	TraceNo       string      `json:"traceNo"`
	TotalFee      string      `json:"totalFee"`
	TransStatus   TransStatus `json:"transStatus"`
	TransDatetime string      `json:"transDatetime"`
	ChannelType   string      `json:"channelType"`
	WebChn        string      `json:"WEB_CHN"`
	Eemark        string      `json:"remark"`
}
