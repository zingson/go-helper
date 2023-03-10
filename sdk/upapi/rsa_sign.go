package upapi

import (
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"sort"
	"strings"
)

// 云闪付 Rsa 私钥签名
func UpRsaSign(params *BodyMap, priKey string, containNilVal bool) (sign string, err error) {
	if priKey == "" {
		return "", errors.New("ERROR:云闪付接口 RSA Sign privateKey 私钥配置不能为空")
	}
	value := rsaSignSortMap(params, containNilVal)
	return RsaSign(value, priKey)
}

func RsaSign(value, priKey string) (sign string, err error) {
	var der []byte
	if strings.HasPrefix(priKey, "-----") {
		p, _ := pem.Decode([]byte(priKey))
		der = p.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(priKey)
		if err != nil {
			return
		}
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return
	}

	hash := sha256.New()
	hash.Write([]byte(value))
	shaBytes := hash.Sum(nil)
	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

// 云闪付 Rsa 公钥验签
func UpRsaVerify(sign string, params *BodyMap, pubKey string, containNilVal bool) (err error) {
	if pubKey == "" {
		return errors.New("ERROR:云闪付接口 RSA Sign upPublicKey 公钥未配置")
	}
	value := rsaSignSortMap(params, containNilVal)
	return RsaVerify(sign, value, pubKey)
}

func RsaVerify(sign, value, pubKey string) (err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		err = errors.New("验签错误，Base64解码签名出错 " + err.Error())
		return
	}
	var der []byte
	if strings.HasPrefix(pubKey, "-----") {
		p, _ := pem.Decode([]byte(pubKey))
		if p == nil {
			err = errors.New("验签错误，PemDecodePublicKey Error")
			return
		}
		der = p.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(pubKey)
		if err != nil {
			return
		}
	}

	publicKey, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		err = errors.New("验签错误，ParsePKIXPublicKey " + err.Error())
		return
	}
	hash := sha256.New()
	hash.Write([]byte(value))
	shaBytes := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, shaBytes, signBytes)
	if err != nil {
		err = errors.New("SIGN_ERROR " + err.Error())
		return
	}
	return
}

// @params containNilVal true空字段参与签名 false空字段不参与签名
func rsaSignSortMap(params *BodyMap, containNilVal bool) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range params.m {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if "signature" == k || "symmetricKey" == k {
			continue
		}
		// 不包含value为空的字段
		if !containNilVal && params.Get(k) == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(params.Get(k))
		buf.WriteByte('&')
	}
	s := buf.String()
	s = s[0 : len(s)-1]
	return s
}
