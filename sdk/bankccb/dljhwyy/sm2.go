package dljhwyy

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

// 2023-10-24

// Sm2Decode 建行-约惠大连微应用-授权参数解密
func Sm2Decode(data string, priKey string) (v string, err error) {
	defer func() {
		if err != nil {
			err = errors.New("ERR_CCB:" + err.Error())
		}
	}()
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}
	pribyte, err := base64.StdEncoding.DecodeString(priKey)
	if err != nil {
		return
	}
	privateKey, err := x509.ParsePKCS8UnecryptedPrivateKey(pribyte)
	if err != nil {
		return
	}
	b, err := sm2.Decrypt(privateKey, ciphertext, sm2.C1C3C2)
	if err != nil {
		return
	}

	v = string(b)
	return
}

// Sm2Encode 建行生活加密
func Sm2Encode(data string, pubKey string) (v string, err error) {
	defer func() {
		if err != nil {
			err = errors.New("ERR_CCB:" + err.Error())
		}
	}()

	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return
	}
	publicKey, err := x509.ParseSm2PublicKey(pubKeyBytes)
	if err != nil {
		return
	}
	b, err := sm2.Encrypt(publicKey, []byte(data), rand.Reader, sm2.C1C3C2)
	if err != nil {
		return
	}
	v = base64.StdEncoding.EncodeToString(b)
	return
}
