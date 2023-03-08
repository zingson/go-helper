package helper

import (
	"encoding/json"
	"fmt"
	douyinpoi "github.com/zingson/goh/sdk/sdk-douyin-poi"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetAccessToken 为了保障应用的数据安全，只能在开发者服务器使用 AppSecret，如果小程序存在泄露 AppSecret 的问题，字节小程序平台将有可能下架该小程序，并暂停该小程序相关服务。
// access_token 是小程序的全局唯一调用凭据，开发者调用小程序支付时需要使用 access_token。access_token 的有效期为 2 个小时，需要定时刷新 access_token，重复获取会导致之前一次获取的 access_token 的有效期缩短为 5 分钟。
// 接口文档地址:https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/get-access-token
func GetAccessToken(conf *douyinpoi.ClientTokenReq) (access_token string, err error) {
	url := GET_ACCESS_TOKEN_URL
	req := &AccessTokenRequest{
		Appid:     conf.AppId,
		Secret:    conf.AppSecret,
		GrantType: "client_credential",
	}
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("json marshal err:%v", err)
	}

	request, err := http.NewRequest("POST", WEBSITE_URl+url, strings.NewReader(string(reqJson)))
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

	accessTokenResult := &AccessTokenResponse{}
	if err = json.Unmarshal(resBytes, accessTokenResult); err != nil {
		return
	}
	if accessTokenResult.ErrNo == 0 {
		access_token = accessTokenResult.Data.AccessToken
	}
	return
}

type AccessTokenRequest struct {
	Appid     string `json:"appid"`      // 小程序 ID
	Secret    string `json:"secret"`     // 小程序的 APP Secret
	GrantType string `json:"grant_type"` // 获取 access_token 时值为 client_credential
}

type AccessTokenResponse struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Expire      int    `json:"expire"`
	} `json:"data"`
}
