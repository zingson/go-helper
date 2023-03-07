package banknbcb

import "encoding/json"

//JsPay 微信公众号/小程序支付
func JsPay(c *Config, params *JsPayParams) (r *JsPayResult, err error) {
	b0, err := json.Marshal(params)
	if err != nil {
		return
	}
	p := make(map[string]interface{})
	err = json.Unmarshal(b0, &p)
	if err != nil {
		return
	}
	p["tran_code"] = "jsPay"
	p["mchnt_cd"] = c.MchntCd

	var wr *JsPayWx
	err = POST(c, p, &wr)
	r = wr.JsPayResult
	r.Package = wr.PackageData
	return
}

type JsPayParams struct {
	ShopId          string `json:"shop_id"`
	SubAppid        string `json:"sub_appid"`
	SubOpenid       string `json:"sub_openid"`
	NotifyUrl       string `json:"notify_url"`
	TraceNo         string `json:"trace_no"`
	TotalFee        string `json:"total_fee"`
	IsCredit        string `json:"isCredit"` // true/false   false:不可 用信用卡支 付
	Remark          string `json:"remark"`
	DuplicateNoFlag string `json:"duplicate_no_flag"` //Y—trace_no 支持重复上送  N—trace_no 不支持重复上 送
}

type JsPayWx struct {
	PackageData string `json:"packageData"` // 微信字段名是package
	*JsPayResult
}

type JsPayResult struct {
	AppId     string `json:"appId"`
	Package   string `json:"package"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	PaySign   string `json:"paySign"`
	SignType  string `json:"signType"`
}

func (o *JsPayResult) String() (s string) {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *JsPayResult) JSON() (s string) {
	b, _ := json.Marshal(o)
	return string(b)
}
