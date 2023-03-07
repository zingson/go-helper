package allinpay

//Config 收银宝接口配置参数
type Config struct {
	BaseUrl string `json:"baseUrl"` // 基础服务地址 https://vsp.allinpay.com
	Cusid   string `json:"cusid"`   // 商户号
	Appid   string `json:"appid"`   // 应用编号
	PriKey  string `json:"priKey"`  // RSA base64格式 商户私钥 PKCS8
	PubKey  string `json:"pubKey"`  // RSA base64格式 通联公钥
}
