package wxapp

import (
	"encoding/json"
)

// GetPhoneNumber
// 解密小程序返回的密文手机号  https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html
// deprecated: 微信改为code方式，使用接口  https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
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
