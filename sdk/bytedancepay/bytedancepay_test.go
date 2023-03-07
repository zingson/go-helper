package bytedancepay

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"testing"
)

//go:embed .secret/fxc.toml
var tomlS string

func Conf() (c *Config) {
	_, err := toml.Decode(tomlS, &c)
	if err != nil {
		panic(err)
	}
	return
}

// 测试创建订单
func TestCreateOrder(t *testing.T) {
	conf := Conf()
	data, err := CreateOrder(conf, &CreateOrderParams{
		AppId:       conf.AppId,
		OutOrderNo:  Rand32(),
		TotalAmount: 1,
		Subject:     "测试商品",
		Body:        "测试商品详情",
		ValidTime:   1800,
		NotifyUrl:   "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
	// 字节小程序      支付 POST https://developer.toutiao.com/api/apps/ecpay/v1/create_order  请求报文      ：{\"app_id\":\"tta549d3f43a70e97401\",\"body\":\"测试商品详情\"       ,\"cp_extra\":\"\",\"disable_msg\":0,\"msg_page\":\"\",\"notify_url\":\"\",\"out_order_no\":\"ac8df697d1b49fd152ee59ac08934011\",\"sign\":\"6751325996cc1a03fd31718de3a61a15\",\"subject\":\" 测试商品\",\"thirdparty_id\":\"\",\"total_amount\":1,\"valid_tim    e\":1800}  响应报文：{\"err_no\":0,\"err_tips\":\"\",\"data\":     {\"order_id\":\"7070057257513421094\",\"order_token\":\"CgwIARDiDRibDiABKAESTgpMQybWpJfQ6ovMBez1XFdaZjchPtiV/jQ6+qdzGHjYEhQBAlz6xihRiJ5NduEWrhj5+cRbwk+WKGPbfqI/fDtrsYWYAurj2LLaoBlcuxoA\"}}    295ms
}

// 测试查询订单
func TestQueryOrder(t *testing.T) {
	conf := Conf()
	rs, err := QueryOrder(conf, &QueryOrderParams{
		AppId:      conf.AppId,
		OutOrderNo: "ac8df697d1b49fd152ee59ac08934011",
	})
	if err != nil {
		return
	}
	t.Log(rs.Json())
}

func TestName(t *testing.T) {
	t.Log(SUCCESS)
}
