package v2

import (
	"encoding/json"
	"testing"
)

var cfg = &Config{
	BaseServiceUrl: "http://mbsmemberwebapi.test.onfishes.com",
	AppKey:         "211394653",
	AppSecret:      "cyQpQKMaGUG186iEZRt1uQ==",
	RsaPriKey:      `MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAOEURCShlw8mfr0xvB6UpFK86Vr2LzEkg6YMNpREvGYmhSbgEIE0KCmH7gIvKr6AYzc78nN7IorW8DusupPIrPSBrHI2b3f0nfFGrAh84mK8kdc9UceNnIlp4caWWLCsbglxiJKcQpz/B3FbHZ9xOboUaT1btByx1ku+4Ri5ZMCLAgMBAAECgYASmOEUgcGAf/bC3SQlBrUZHQDPAj5d+h1ij+nGkHNcoVwpHSRf+JZE8DVLOuh2Oxd3jd13izoMbLwGwjvcUB1yIsM35lgJRw4Ofj4eqK5peP94LpTPETnufyM2BveIFgUhS7P0Qt7K/5o/xjdzj87ebsTxzxnUUU410yPYisN1wQJBAPWS84yn7yWZQTo06jDCkKI52Elq66Z6NSRWzOC0CgTYyu7VzN29URrIxT+7353lZFM2xlA6syvpMiVYCqdGkaECQQDqopBplN9R0MCs/GTv7r2pBpP27YqiNKNcfE/iq71qZcOZwdrupSqvYOHil7OpzPjf0G8x4SnV/KXN1ppPYzqrAkEArqBX1h6ZHWh0jLqSCihhBysRFWwVtGVUosmimOsN8NJkxB9+tfNo2B4Kvb6QTkyP4eiibuy++iuygAGyWa8B4QJBAOlJCFxxcDhgXbGgoJsNu/S6XZM9SoFL5MCnuKWeK44F8ByH6a0s+uu0X+JzAmbpLOkay/PD81yW/iNSI8qa1lECQE5GBYFkZlvLgoVoNrEZval8az2FLQgyzyWcbypTruvWCD4LQSAXdY8Wc5jWTMT5m4IyG1UeqX8DrIRquoYXvdY=`,
	RsaPubKey:      `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDhFEQkoZcPJn69MbwelKRSvOla9i8xJIOmDDaURLxmJoUm4BCBNCgph+4CLyq+gGM3O/JzeyKK1vA7rLqTyKz0gaxyNm939J3xRqwIfOJivJHXPVHHjZyJaeHGlliwrG4JcYiSnEKc/wdxWx2fcTm6FGk9W7QcsdZLvuEYuWTAiwIDAQAB`,
}

var c = NewClient(cfg)

func TestGetAccount(t *testing.T) {
	// 查询账户余额
	r, err := c.GetAccount()
	if err != nil {
		t.Error(err)
		return
	}
	bytes, _ := json.Marshal(r)
	t.Log(string(bytes))
}

// 订单查询
func TestOrderQueryV2(t *testing.T) {
	r, err := c.OrderQueryV2("990720101059396205")
	if err != nil {
		t.Error(err)
		return
	}
	bytes, _ := json.Marshal(r)
	t.Log(string(bytes))
}

// 订单提交接口
func TestOrderInsertV2(t *testing.T) {
	r, err := c.OrderInsertV2(&OrderInsertV2Params{
		MemberAmountCode: "",
		ProductCode:      "PLM100068",
		BuyCount:         1,
		MOrderID:         "890720101059396205",
		ChargeAccount:    "13611703040",
		ExtendParam:      nil,
		CallBackUrl:      "https://msd.himkt.cn/voucher/yushang/callback/notify.do",
	})
	if err != nil {
		t.Error(err)
		return
	}
	bytes, _ := json.Marshal(r)
	t.Log(string(bytes))
}

// 卡密申请接口
func TestOrderInsertSiberianNitrariaFruitV2(t *testing.T) {
	r, err := c.OrderInsertSiberianNitrariaFruitV2(&OrderInsertSiberianNitrariaFruitV2Params{
		MemberAmountCode: "",
		ProductCode:      "PLM100010",
		BuyCount:         1,
		MOrderID:         "990720101059396205",
		ChargeAccount:    "",
		CallBackUrl:      "https://msd.himkt.cn/voucher/onfishes/callback.do",
	})
	if err != nil {
		t.Error(err)
		return
	}
	bytes, _ := json.Marshal(r)
	t.Log(string(bytes))

}

func TestDecode(t *testing.T) {
	text := "lVnMSzfi1+gsawtb6zrWmOImWdAkpgEJ00Qsx0x/GVvT5yR6Wr4aQameWfw0bGhjcTJ7xg0tVLT/yz9ZnYyCSrAmWwa71gO+cBTj7ZyD0mmrUBNLc4vh/pN+biRHiYuoevKI+E0yZLv6KY1/uIVZWy0WVVJyb66Wx3pKzPtr2xY="
	val, err := RsaDecrypt(text, cfg.RsaPriKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}
