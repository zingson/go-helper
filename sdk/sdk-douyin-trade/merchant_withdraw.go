package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// MerchantWithdraw 商户提现
// 接口文档地址 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/withdraw/merchant-withdraw
func MerchantWithdraw(req *MerchantWithdrawRequest) (res string, err error) {
	url := MERCHANT_WITHDRAW_URL

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
	println("MerchantWithdraw,request:%s, response:%s", string(rBytes), res)
	return
}

type MerchantWithdrawRequest struct {
	ThirdpartyId   string `json:"thirdparty_id,omitempty"`
	AppId          string `json:"app_id,omitempty"`
	MerchantUid    string `json:"merchant_uid"`
	ChannelType    string `json:"channel_type"`
	WithdrawAmount int64  `json:"withdraw_amount"`
	OutOrderId     string `json:"out_order_id"`
	Callback       string `json:"callback,omitempty"`
	CpExtra        string `json:"cp_extra,omitempty"`
	Sign           string `json:"sign"`
}
