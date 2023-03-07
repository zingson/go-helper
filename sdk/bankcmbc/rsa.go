package bankcmbc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"strings"
)

func der(key string) (derBytes []byte) {
	if strings.Contains(key, "-----") {
		p, _ := pem.Decode([]byte(key))
		if p == nil {
			panic(errors.New("Key pem.Decode error "))
		}
		derBytes = p.Bytes
	} else {
		bytes, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			panic(errors.New("pubKey Base64解码错误 " + err.Error()))
		}
		derBytes = bytes
	}
	return
}

//RsaSign 签名
func RsaSign(msg, priKey string) (sign string, err error) {
	msg = strings.TrimSpace(msg)
	msg = strings.Trim(msg, "\n")
	pri, err := x509.ParsePKCS1PrivateKey(der(priKey))
	if err != nil {
		return
	}
	hash := sha1.New()
	hash.Write([]byte(msg))
	hashed := hash.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA1, hashed)
	if err != nil {
		return
	}
	sign = hex.EncodeToString(sig)
	return
}

//RsaVery 验签
func RsaVery(msg, sign, pubKey string) (err error) {
	msg = strings.TrimSpace(msg)
	msg = strings.Trim(msg, "\n")
	pub, err := x509.ParsePKCS1PublicKey(der(pubKey))
	if err != nil {
		return
	}
	hash := sha1.New()
	hash.Write([]byte(msg))
	hashed := hash.Sum(nil)
	sig, err := hex.DecodeString(sign)
	if err != nil {
		return
	}
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed, sig)
}
