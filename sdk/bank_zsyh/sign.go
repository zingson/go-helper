package bank_zsyh

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
