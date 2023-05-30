package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

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
	b, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
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
