package bytedancepay

type Config struct {
	ServiceUrl string `json:"service_url"` // 接口域名
	AppId      string `json:"app_id"`      // 小程序APPID
	Mchid      string `json:"mchid"`       // 支付商户号
	Salt       string `json:"salt"`        // 支付Salt
	Cburl      string `json:"cburl"`       // 支付回调URL
	Token      string `json:"token"`       // Token令牌，小程序支付配置
}
