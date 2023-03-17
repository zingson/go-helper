package ccblife_pay

import "encoding/json"

// 建行生活支付通知  请求是POST  参数在URI上

// path=/ccb-pay-notify/105004453993206?POSID=062498142&BRANCHID=331000000&ORDERID=14920688136954019840&PAYMENT=0.01&CURCODE=01&REMARK1=&REMARK2=YS44000007000457**LZF&ACC_TYPE=02&SUCCESS=Y&SIGN=32ea4df898745b02bbc7ce9eca409772cb1c2f45beda4ad520f371861ab9f071bca481db95612afeaa5759ca7e4937c3ca5eea48c91ebc494f22651c245f88e68720e72a9ecb130efacf03617c9391acc7d7130a5cd5abf58217f01d9c84a1a631d923a7533413e59ac38233bcae8405797e081a8eef4c658a5c441d4c922e3e
type PayNotifyParams struct {
	POSID    string `json:"POSID"`
	BRANCHID string `json:"BRANCHID"`
	ORDERID  string `json:"ORDERID"`
	PAYMENT  string `json:"PAYMENT"`
	CURCODE  string `json:"CURCODE"`
	ISECNY   string `json:"ISECNY"` //统一支付网关，如果客户使用的是数币支付，服务器通知及页面通知会增加此字段 ISECNY=Y
	REMARK1  string `json:"REMARK1"`
	REMARK2  string `json:"REMARK2"`
	ACC_TYPE string `json:"ACC_TYPE"`
	SUCCESS  string `json:"SUCCESS"` // Y || N  ，Y=成功 其它不成功
	SIGN     string `json:"SIGN"`
}

func (p *PayNotifyParams) Json() string {
	b, _ := json.Marshal(p)
	return string(b)
}
