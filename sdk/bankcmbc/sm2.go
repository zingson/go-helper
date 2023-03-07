package bankcmbc

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/x509"
)

// Sm2Decode 解密
func Sm2Decode(v string, priKey string) (s string, err error) {
	ciphertext, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return
	}
	pri, err := x509.ReadPrivateKeyFromHex(priKey)
	if err != nil {
		return
	}
	pBytes, err := pri.DecryptAsn1(ciphertext)
	if err != nil {
		return
	}
	s = string(pBytes)
	return
}

// Sm2Encode 加密
func Sm2Encode(v string, pubKey string) (s string, err error) {
	if err != nil {
		return
	}
	pub, err := x509.ReadPublicKeyFromHex(pubKey)
	if err != nil {
		return
	}
	pBytes, err := pub.EncryptAsn1([]byte(v), rand.Reader)
	if err != nil {
		return
	}
	s = base64.StdEncoding.EncodeToString(pBytes)
	return
}

//Sm2Sign 签名
func Sm2Sign(msg, priKey string) (sign string, err error) {
	msg = hex.EncodeToString([]byte(msg))
	pri, err := x509.ReadPrivateKeyFromHex(priKey)
	if err != nil {
		return
	}
	b, err := pri.Sign(rand.Reader, []byte(msg), nil)
	if err != nil {
		return
	}
	sign = hex.EncodeToString(b)
	return
}

// Sm2Very 验签
func Sm2Very(msg, sign, pubKey string) (ok bool, err error) {
	msg = hex.EncodeToString([]byte(msg))
	b, err := hex.DecodeString(sign)
	if err != nil {
		return
	}
	pub, err := x509.ReadPublicKeyFromHex(pubKey)
	if err != nil {
		return
	}
	ok = pub.Verify([]byte(msg), b)
	return
}
