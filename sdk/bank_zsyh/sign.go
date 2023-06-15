package bank_zsyh

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
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

// SortMap map排序
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

func StructToMap(in interface{}) (pmap map[string]string) {
	b, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &pmap)
	if err != nil {
		panic(err)
	}
	return
}

// RsaVerify 验签
func RsaVerify(sign, value, pubKey string) (err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		err = errors.New("验签错误，Base64解码出错 " + err.Error())
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
	hash := sha1.New()
	hash.Write([]byte(value))
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), signBytes)
	if err != nil {
		err = errors.New("验签错误，" + err.Error())
		return
	}
	return
}
