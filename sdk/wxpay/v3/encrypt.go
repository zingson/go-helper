package v3

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// 微信APIv3 签名验签、回调数据解密、敏感数据加密解密

/*
声明所使用的证书:
某些情况下，将需要更新密钥对和证书。为了保证更换过程中不影响API的使用，请求和应答的HTTP头部中包括证书序列号，以声明签名或者加密所用的密钥对和证书。
商户签名使用商户私钥，证书序列号包含在请求HTTP头部的Authorization的serial_no
微信支付签名使用微信支付平台私钥，证书序列号包含在应答HTTP头部的Wechatpay-Serial
商户上送敏感信息时使用微信支付平台公钥加密，证书序列号包含在请求HTTP头部的Wechatpay-Serial
*/

// PrivateKeyPemParse 私钥pem格式解析
func PrivateKeyPemParse(priPem string) (pri *rsa.PrivateKey, err error) {
	keyBytes := []byte(priPem)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		err = errors.New("wxpay rsa private key decode error")
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		err = errors.New("wxpay rsa private key x509 parse error" + err.Error())
		return
	}
	pri = privateKey.(*rsa.PrivateKey)
	return
}

// CertificateParse 微信平台验签公钥pem格式解析
func CertificateParse(pubPem string) (pub *rsa.PublicKey, err error) {
	block, _ := pem.Decode([]byte(pubPem))
	if block == nil {
		err = errors.New("wxpay rsa public key decode error")
		return
	}
	pk, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = errors.New("wxpay rsa public key x509 parse error " + err.Error())
		return
	}

	pub = pk.PublicKey.(*rsa.PublicKey)
	return
}

// RsaSignWithSha256 私钥签名
func RsaSignWithSha256(data string, priKey *rsa.PrivateKey) (string, error) {
	dataBytes := []byte(data)
	h := sha256.New()
	h.Write(dataBytes)
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// RsaVeryWithSha256 公钥验签
func RsaVeryWithSha256(data, signature string, pubKey *rsa.PublicKey) (bool, error) {
	oldSign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	hashed := sha256.Sum256([]byte(data))
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], oldSign)
	if err != nil {
		return false, err
	}
	return true, nil
}

// AesGcmDecrypt 证书和回调报文解密
func AesGcmDecrypt(ciphertext string, nonce, additionalData string, v3Secret string) (plaintext string, err error) {
	key := []byte(v3Secret) //key是APIv3密钥，长度32位，由管理员在商户平台上自行设置的
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	aesGcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return
	}
	cipherData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	plainData, err := aesGcm.Open(nil, []byte(nonce), cipherData, []byte(additionalData))
	if err != nil {
		return
	}
	plaintext = string(plainData)
	return
}

// RsaEncrypt 敏感信息加密
func RsaEncrypt(plaintext []byte, pub *rsa.PublicKey) (ciphertext string, err error) {
	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, plaintext, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}
	ciphertext = base64.StdEncoding.EncodeToString(cipherdata)
	return
}

// RsaDecrypt 敏感信息解密
func RsaDecrypt(ciphertext string, priv *rsa.PrivateKey) (plaintext []byte, err error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	plaintext, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, cipherdata, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}
	return
}
