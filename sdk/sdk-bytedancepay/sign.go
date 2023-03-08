package bytedancepay

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

//getSign 请求签名算法
func getSign(paramsMap map[string]interface{}, secret string) string {
	var paramsArr []string
	for k, v := range paramsMap {
		if k == "other_settle_params" {
			continue
		}
		value := strings.TrimSpace(fmt.Sprintf("%v", v))
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") && len(value) > 1 {
			value = value[1 : len(value)-1]
		}
		value = strings.TrimSpace(value)
		if value == "" || value == "null" {
			continue
		}
		switch k {
		// app_id, thirdparty_id, sign 字段用于标识身份，不参与签名
		case "app_id", "thirdparty_id", "sign":
		default:
			paramsArr = append(paramsArr, value)
		}
	}
	paramsArr = append(paramsArr, secret)
	sort.Strings(paramsArr)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(paramsArr, "&"))))
}

//CbSign 回调签名算法
func CbSign(p map[string]interface{}, token string) string {
	sortedString := make([]string, len(p)+1)
	sortedString = append(sortedString, token)
	for _, v := range p {
		value, _ := v.(string)
		sortedString = append(sortedString, value)
	}
	sort.Strings(sortedString)
	h := sha1.New()
	h.Write([]byte(strings.Join(sortedString, "")))
	bs := h.Sum(nil)
	_signature := fmt.Sprintf("%x", bs)
	return _signature
}

//CbSignNew 回调签名算法
func CbSignNew(token, timestamp, nonce, msg string) string {
	sortedString := make([]string, 4)
	sortedString = append(sortedString, token)
	sortedString = append(sortedString, timestamp)
	sortedString = append(sortedString, nonce)
	sortedString = append(sortedString, msg)
	sort.Strings(sortedString)
	h := sha1.New()
	h.Write([]byte(strings.Join(sortedString, "")))
	bs := h.Sum(nil)
	_signature := fmt.Sprintf("%x", bs)
	return _signature
}
