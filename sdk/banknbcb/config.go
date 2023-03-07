package banknbcb

type Config struct {
	ServiceUrl string `json:"serviceUrl"` // 服务地址
	MchntCd    string `json:"mchntCd"`    // 易收宝分配的 商户号
	MchPriKey  string `json:"mchPriKey"`  // 商户私钥，易收宝提供
	NbcbPubKey string `json:"nbcbPubKey"` // 易收宝公钥，易收宝提供
}
