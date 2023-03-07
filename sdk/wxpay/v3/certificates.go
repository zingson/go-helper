package v3

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"time"
)

/*
文档：https://wechatpay-api.gitbook.io/wechatpay-api-v3/jie-kou-wen-dang/ping-tai-zheng-shu

获取商户当前可用的平台证书列表。微信支付提供该接口，帮助商户后台系统实现平台证书的平滑更换。

注意事项
如果自行实现验证平台签名逻辑的话，需要注意以下事项:
程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
定期调用该接口，间隔时间小于12 小时
加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）

说明：
在启用新的平台证书前，微信支付会提前24小时把新证书加入到平台证书列表中
接口的频率限制: 单个商户号1000 次/s
首次下载证书，可以使用微信支付提供的证书下载工具
*/

// Certificates 微信平台证书列表, 返回最新证书
func Certificates(c *Client) (cert *Cert, err error) {
	var (
		method = "GET"
		path   = "/v3/certificates"
	)

	rBytes, _, err := Do(c, method, path, "")
	if err != nil {
		return
	}

	var result *CertificatesResult
	err = json.Unmarshal(rBytes, &result)
	if err != nil {
		return
	}

	cert = new(Cert)
	for _, datum := range result.Data {
		if time.Time(cert.EffectiveTime).Before(time.Time(datum.EffectiveTime)) {
			ce := datum.EncryptCertificate
			ciphertext := ce.Ciphertext
			nonce := ce.Nonce
			associatedData := ce.AssociatedData
			plaintext, err := AesGcmDecrypt(ciphertext, nonce, associatedData, c.V3Secret)
			if err != nil {
				fmt.Println("Error AesGcmDecrypt：", err.Error())
				return nil, err
			}
			pub, err := CertificateParse(plaintext)
			if err != nil {
				fmt.Println("Error PublicKeyPemParse：", err.Error())
				return nil, err
			}
			cert.PublicKey = pub
			cert.EffectiveTime = datum.EffectiveTime
			cert.ExpireTime = datum.ExpireTime
			cert.SerialNo = datum.SerialNo
		}
	}
	return
}

type Cert struct {
	PublicKey     *rsa.PublicKey
	SerialNo      string
	EffectiveTime Time
	ExpireTime    Time
}

type CertificatesResult struct {
	Data []*CertificatesResultData `json:"data"`
}

type CertificatesResultData struct {
	SerialNo           string                         `json:"serial_no"`
	EffectiveTime      Time                           `json:"effective_time"`
	ExpireTime         Time                           `json:"expire_time"`
	EncryptCertificate *CertificatesResultDataEncrypt `json:"encrypt_certificate"`
}

type CertificatesResultDataEncrypt struct {
	Algorithm      string `json:"algorithm"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
}
