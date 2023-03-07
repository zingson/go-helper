package v3

import (
	"crypto/rsa"
)

// V3Sign 计算请求报文签名
func V3Sign(method, path, body, timestamp, nonceStr string, priKey *rsa.PrivateKey) (sign string, err error) {
	targetStr := method + "\n" + path + "\n" + timestamp + "\n" + nonceStr + "\n" + body + "\n"
	return RsaSignWithSha256(targetStr, priKey)
}

// V3SignVery 验签响应报文签名
func V3SignVery(signature, time, nonce, body string, pubKey *rsa.PublicKey) (ok bool, err error) {
	data := time + "\n" + nonce + "\n" + body + "\n"
	return RsaVeryWithSha256(data, signature, pubKey)
}
