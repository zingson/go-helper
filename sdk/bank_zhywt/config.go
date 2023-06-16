package bank_zhywt

import "sync"

//招行一网通支付
//商户服务平台 https://pay.cmbchina.com
//商户服务平台获取的商户号  前4位是分行号 后6位是商户号

// Config 招行一网通接口配置参数
type Config struct {
	MerchservUrl  string `json:"merchserv_url"`  //https://merchserv.netpay.cmbchina.com
	CmbBankB2BUrl string `json:"cmbbankb2b_url"` //https://b2b.cmbchina.com
	NetpaymentUrl string `json:"netpayment_url"` //https://netpay.cmbchina.com
	BranchNo      string `json:"branch_no"`      //商户分行号，4位数字	0755
	MerchantNo    string `json:"merchant_no"`    //商户号，6位数字	058624
	Merkey        string `json:"mer_key"`        //商户秘钥
}

// 缓存 用Base64编码的招行公钥
var fbPubKeyMap sync.Map

func (c *Config) GetFbPubKey() (fbPubKey string, err error) {
	key := c.BranchNo + c.MerchantNo
	if v, ok := fbPubKeyMap.Load(key); ok {
		fbPubKey = v.(string)
		return
	}
	fbPubKey, err = QueryKeyAPI(c)
	if err != nil {
		return
	}
	fbPubKeyMap.Store(key, fbPubKey)
	return
}
