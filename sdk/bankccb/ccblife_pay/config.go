package ccblife_pay

type Config struct {
	ServiceSvcjson    string `json:"service_svcjson" toml:"service_svcjson"`       // 服务地址
	ServiceOccplatreq string `json:"service_occplatreq" toml:"service_occplatreq"` // 退款和订单查询 服务地址
	MerchantId        string `json:"merchant_id" toml:"merchant_id"`               // 商户号
	PosId             string `json:"pos_id" toml:"pos_id"`                         // 商户柜台代码
	BranchId          string `json:"branch_id" toml:"branch_id"`                   // 分行代码
	PlatformId        string `json:"platform_id" toml:"platform_id"`               // 服务方编号
	PubKey            string `json:"pub_key" toml:"pub_key"`                       // 服务方公钥
	PriKey            string `json:"pri_key" toml:"pri_key"`                       // 服务方私钥 解密用户授权信息
	MchPubKey         string `json:"mch_pub_key" toml:"mch_pub_key"`               // 建行生活商户公钥，私钥建行生活保留
}
