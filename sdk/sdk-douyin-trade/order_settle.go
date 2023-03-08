package helper

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// OrderSettle 担保交易的发起结算及分账:到达分账周期后，商户将交易成功的资金，结算给自己或分账给其他分账方。
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/settlements/settlement
func OrderSettle(req *OrderSettleRequest) (res string, err error) {
	url := ORDER_SETTLE_URL

	pmap := structToMap(req)
	pmap["sign"] = getSign(pmap, SALT)
	rBytes, err := json.Marshal(pmap)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}

	request, err := http.NewRequest("POST", WEBSITE_URl+url, strings.NewReader(string(rBytes)))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	defer request.Body.Close()
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	res = string(resBytes)
	println("000000000000000:%s", res)
	return
}

type OrderSettleRequest struct {
	AppId       string `json:"app_id"`
	OutSettleNo string `json:"out_settle_no"`
	OutOrderNo  string `json:"out_order_no"`
	SettleDesc  string `json:"settle_desc"`
	//SettleParams string `json:"settle_params,omitempty"`
	Sign string `json:"sign"`
	//CpExtra      string `json:"cp_extra,omitempty"`
	//NotifyUrl    string `json:"notify_url,omitempty"`
	//ThirdpartyId string `json:"thirdparty_id,omitempty"`
	//Finish       string `json:"finish,omitempty"`
}

func structToMap(in interface{}) (pmap map[string]interface{}) {
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
