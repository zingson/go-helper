package bytedanceapp

import (
	"encoding/json"
)

/*
注意：
使用（用户点击 button）前需先调用tt.login接口。如果在回调中调用 tt.login 会刷新登录态，导致登录后换取的 session_key 与手机号码加密时使用的 session_key 不同，从而导致解密失败。
*/

//GetPhoneNumber 解密小程序返回的密文手机号  https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/component/acquire-phone-number-acquire/
func GetPhoneNumber(appid, sessionKey, iv, encryptedData string) (r *PhoneNumberResult, err error) {
	rbytes, err := DecryptData(appid, sessionKey, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(rbytes, &r)
	if err != nil {
		return
	}
	return
}

type PhoneNumberResult struct {
	PhoneNumber     string `json:"phoneNumber"`     //用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string `json:"purePhoneNumber"` //没有区号的手机号
	CountryCode     string `json:"countryCode"`     //区号
}

func (o *PhoneNumberResult) JSON() string {
	pbs, _ := json.Marshal(o)
	return string(pbs)
}
