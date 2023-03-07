package ccblife_pay

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

const (
	MAX_ENCRYPT_BLOCK = 117 //RSA最大加密明文大小
	MAX_DECRYPT_BLOCK = 128 //RSA最大解密密文大小
)

//RsaDecode 建行生活解密
func RsaDecode(data string, priKey string) (v string, err error) {
	defer func() {
		if err != nil {
			err = errors.New("ERR_CCB:" + err.Error())
		}
	}()

	// 数据Base64解码
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		err = errors.New("1-Base64解码错误 " + err.Error())
		return
	}
	ciphertext, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		err = errors.New("2-Base64解码错误 " + err.Error())
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

	// RSA 分段解密
	buffer := bytes.Buffer{}
	var offset = 0
	var inputLen = len(ciphertext)
	for inputLen-offset > 0 {
		var offsetLen int
		if inputLen-offset > MAX_DECRYPT_BLOCK {
			offsetLen = offset + MAX_DECRYPT_BLOCK
		} else {
			offsetLen = inputLen
		}
		var cache []byte
		cache, err = rsa.DecryptPKCS1v15(rand.Reader, pri.(*rsa.PrivateKey), ciphertext[offset:offsetLen])
		if err != nil {
			err = errors.New("rsa.DecryptPKCS1v15 error " + err.Error())
			return
		}
		buffer.Write(cache)
		offset = offset + MAX_DECRYPT_BLOCK
	}

	// RSA解密结果
	v = string(buffer.Bytes())
	return
}

//RsaEncode 建行生活加密
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

	// RSA 分段解密
	buffer := bytes.Buffer{}
	var offset = 0
	var inputLen = len(dBytes)
	for inputLen-offset > 0 {
		var offsetLen int
		if inputLen-offset > MAX_ENCRYPT_BLOCK {
			offsetLen = offset + MAX_ENCRYPT_BLOCK
		} else {
			offsetLen = inputLen
		}
		var cache []byte
		cache, err = rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), dBytes[offset:offsetLen])
		if err != nil {
			err = errors.New("rsa.EncryptPKCS1v15 error " + err.Error())
			return
		}
		buffer.Write(cache)
		offset = offset + MAX_ENCRYPT_BLOCK
	}

	// RSA解密结果 Base64解码
	v = base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString(buffer.Bytes())))
	return
}
