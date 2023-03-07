package ldpush

type Config struct {
	ServiceUrl string `json:"service_url"`
	Appid      string `json:"appid"`
	RsaPriKey  string `json:"rsa_pri_key"` // 接入私钥，公钥提供给平台
}
