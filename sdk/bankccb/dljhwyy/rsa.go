package dljhwyy

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

// RsaDecode 建行-约惠大连微应用-授权参数解密
func RsaDecode(data string, priKey string) (v string, err error) {
	defer func() {
		if err != nil {
			err = errors.New("ERR_CCB:" + err.Error())
		}
	}()

	// 数据Base64解码
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		err = errors.New("Base64解码错误 " + err.Error())
		return
	}

	// RSA  公钥支持pem格式与base格式
	var der []byte
	if strings.Contains(priKey, "BEGIN PRIVATE KEY") {
		p, _ := pem.Decode([]byte(priKey))
		if p == nil {
			err = errors.New("priKey pem.Decode error ")
			return
		}
		der = p.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(priKey)
		if err != nil {
			err = errors.New("priKey Base64解码错误 " + err.Error())
			return
		}
	}

	// RSA 密钥解析
	pri, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		err = errors.New("x509.ParsePKCS8PrivateKey error " + err.Error())
		return
	}

	// RSA 解密
	rsBytes, err := rsa.DecryptPKCS1v15(rand.Reader, pri.(*rsa.PrivateKey), ciphertext)
	if err != nil {
		err = errors.New("rsa.DecryptPKCS1v15 error " + err.Error())
		return
	}

	// RSA解密结果
	v = string(rsBytes)
	return
}

// RsaEncode 建行生活加密
func RsaEncode(data string, pubKey string) (v string, err error) {
	defer func() {
		if err != nil {
			err = errors.New("ERR_CCB:" + err.Error())
		}
	}()

	var dBytes = []byte(data)

	// RSA  公钥支持pem格式与base格式
	var der []byte
	if strings.Contains(pubKey, "BEGIN PUBLIC KEY") {
		p, _ := pem.Decode([]byte(pubKey))
		if p == nil {
			err = errors.New("pubKey pem.Decode error ")
			return
		}
		der = p.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(pubKey)
		if err != nil {
			err = errors.New("pubKey Base64解码错误 " + err.Error())
			return
		}
	}

	// RSA 解析加密
	pub, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		err = errors.New("x509.ParsePKIXPublicKey error " + err.Error())
		return
	}

	// RSA 解密
	sBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), dBytes)
	if err != nil {
		err = errors.New("rsa.EncryptPKCS1v15 error " + err.Error())
		return
	}

	// RSA解密结果 Base64解码
	v = base64.StdEncoding.EncodeToString(sBytes)
	return
}
