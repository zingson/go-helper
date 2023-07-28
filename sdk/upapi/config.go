package upapi

// Config 银联云闪付接口配置
type Config struct {
	MiniAppid     string `json:"mini_appid" description:"小程序应用ID"`
	ServiceUrl    string `json:"service_url"  description:"接口链接前缀，结尾不带'/'"` // https://open.95516.com/open/access/1.0
	Appid         string `json:"appid,omitempty" description:"接入方唯一标识"`
	Secret        string `json:"secret,omitempty" description:"接入方秘钥，用于基础令牌接口的签名"`
	SymmetricKey  string `json:"symmetric_key,omitempty" description:"对称密钥（3DES，16进制格式） 。用于后台敏感数据解密"`
	UpPublicKey   string `json:"up_public_key,omitempty" description:"（银联方）使用openssl生成，base64形式输出"`
	MchPrivateKey string `json:"mch_private_key" description:"接入商户证书私钥"`
	AcctEntityTp  string `json:"acct_entity_tp"` //默认03 营销活动配置的赠送维度（参见营销平台活动配置），2位， 可选：01=手机 02-卡号 03-用户（二选一） 赠送维度为卡号时，则cardNo必填； 赠送维度为用户时，则openId，mobile, cardNo三选一上送
}
