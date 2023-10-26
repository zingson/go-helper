package v3

import (
	"crypto/rsa"
	"errors"
	"sync"
)

// Option 微信支付参数配置项
type Option struct {
	ServiceUrl string `json:"service_url" toml:"service_url"` // 可选，微信接口服务地址，默认：https://api.mch.weixin.qq.com
	Mchid      string `json:"mchid" toml:"mchid"`             // 微信支付商户号
	V3Secret   string `json:"v3_secret" toml:"v3_secret"`     // 微信V3秘钥 ，用于证书与回调报文解密 AEAD_AES_256_GCM
	SerialNo   string `json:"serial_no" toml:"serial_no"`     // 微信商户证书序列号
	PrivateKey string `json:"private_key" toml:"private_key"` // 微信支付商户私钥,PEM格式
}

func NewClient(option Option) (c *Client, err error) {
	c = &Client{}
	err = c.SetOption(option)
	return
}

type Client struct {
	Option
	mchPriKey  *rsa.PrivateKey // 微信支付商户私钥
	wxSerialNo string          // 微信平台证书序号,
}

// SetOption 修改Client配置
func (c *Client) SetOption(option Option) (err error) {
	if option.ServiceUrl == "" {
		option.ServiceUrl = "https://api.mch.weixin.qq.com"
	}
	if option.Mchid == "" {
		err = errors.New("ERR_OPTION:微信支付配置缺少商户号")
		return
	}
	if option.PrivateKey == "" || option.SerialNo == "" {
		err = errors.New("ERR_OPTION:微信支付配置缺少商户证书序号与私钥")
		return
	}
	if option.V3Secret == "" {
		err = errors.New("ERR_OPTION:微信支付配置缺少v3Secret")
		return
	}
	mchPriKey, err := PrivateKeyPemParse(option.PrivateKey)
	if err != nil {
		return
	}
	c.mchPriKey = mchPriKey
	c.Option = option
	return
}

// 微信平台证书序号与公钥，使用接口自动更新 ,k=序号 v=证书
var wpk = sync.Map{}

// GetWxPubKey 微信平台公钥
func (c *Client) GetWxPubKey(wxSerial string) (pub *rsa.PublicKey, err error) {
	value, ok := wpk.Load(wxSerial)
	if ok {
		c.wxSerialNo = wxSerial
		pub = value.(*rsa.PublicKey)
		return
	}
	cert, err := Certificates(c)
	if err != nil {
		return
	}
	c.wxSerialNo = cert.SerialNo
	wpk.Store(cert.SerialNo, cert.PublicKey)
	pub = cert.PublicKey
	return
}

func (c *Client) HttpPost(path string, i interface{}, o interface{}) (err error) {
	return Call(c, "POST", path, i, o)
}

func (c *Client) HttpGet(path string, o interface{}) (err error) {
	return Call(c, "GET", path, nil, o)
}
