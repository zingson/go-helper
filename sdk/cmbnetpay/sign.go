package cmbnetpay

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"sort"
	"strings"
)

// Sha256Sign 请求报文签名
func Sha256Sign(targetStr string) (sign string) {
	h := sha256.New()
	h.Write([]byte(targetStr))
	return hex.EncodeToString(h.Sum(nil))
}

// RSAVerify 验签响应报文签名
func RSAVerify(origdata, sign string, publicKey []byte) (bool, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(origdata))
	digest := h.Sum(nil)
	body, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, digest, body)
	if err != nil {
		return false, err
	}
	return true, nil
}

//SortMap map排序
// @params containNilVal true空字段参与签名 false空字段不参与签名
func SortMap(m map[string]string, containNilVal bool) string {
	if m == nil {
		return ""
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range m {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if "sign" == k {
			continue
		}
		// 不包含value为空的字段
		if !containNilVal && m[k] == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(m[k])
		buf.WriteByte('&')
	}
	s := buf.String()
	s = s[0 : len(s)-1]
	return s
}
