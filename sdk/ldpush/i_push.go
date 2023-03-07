package ldpush

import "fmt"

//IPush 语音推送交易   文档：https://artisan.landicorp.com/doc/api?id=839826b01854493699202aeab0029bb2
func IPush(conf *Config, params *PushParams) (err error) {
	pm := map[string]string{
		"msgId":    params.MsgId,
		"sn":       params.Sn,
		"content":  params.Content,
		"totalFee": params.TotalFee,
	}
	if params.TransactionId != "" {
		pm["transactionId"] = params.TransactionId
	}
	if params.PayType != "" {
		pm["payType"] = params.PayType
	}
	if params.DiscountFee != "" {
		pm["discountFee"] = params.DiscountFee
	}
	if params.Attach != "" {
		pm["attach"] = params.Attach
	}
	if params.Lang != "" {
		pm["lang"] = params.Lang
	}

	rbytes, err := Post(conf, "/push", pm)
	if err != nil {
		return
	}
	fmt.Println(rbytes)
	return
}

type PushParams struct {
	MsgId         string // 非空
	Sn            string // 非空
	TransactionId string
	Content       string // 非空
	TotalFee      string // 非空 格式为xxx.xx精确到分
	PayType       string
	DiscountFee   string
	Attach        string
	Lang          string
}
