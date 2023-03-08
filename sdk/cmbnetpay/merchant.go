package cmbnetpay

// FixedParams 接口固定参数，所有接口都包含这部分内容
type FixedParams struct {
	Version  string `json:"version" ` // 接口版本号,固定为”1.0”
	Charset  string `json:"charset"`  // 参数编码,固定为“UTF-8”
	Sign     string `json:"sign"`     // 报文签名,使用商户支付密钥对reqData内的数据进行签名
	SignType string `json:"signType"` // 签名算法,固定为”SHA-256”
}

func GetFixedParams() FixedParams {
	return FixedParams{
		Version:  "1.0",
		Charset:  "UTF-8",
		SignType: "SHA-256",
	}
}
