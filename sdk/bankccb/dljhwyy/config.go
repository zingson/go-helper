package dljhwyy

type Config struct {
	PrivateKey string `json:"private_key"` // 私钥PEM格式，公钥提供给约惠
	Remark     string `json:"remark"`      // 备注
}
