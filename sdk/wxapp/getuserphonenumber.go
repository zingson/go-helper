package wxapp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// GetUserPhoneNumber 根据小程序code获取手机号
// 接口：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func GetUserPhoneNumber(access_token, code string) (phoneInfo PhoneInfo, err error) {
	body := map[string]string{"code": code}
	b, _ := json.Marshal(body)
	resp, err := http.Post("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token="+access_token, "application/json", bytes.NewReader(b))
	if err != nil {
		return
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rs GetUserPhoneNumberResult
	err = json.Unmarshal(respBytes, &rs)
	if err != nil {
		return
	}
	if rs.Errcode != 0 {
		err = errors.New(strconv.FormatInt(rs.Errcode, 10) + ":" + rs.Errmsg)
		return
	}
	phoneInfo = rs.PhoneInfo
	return
}

type GetUserPhoneNumberResult struct {
	Errcode   int64     `json:"errcode"` // 0=成功，其它失败
	Errmsg    string    `json:"errmsg"`
	PhoneInfo PhoneInfo `json:"phone_info"`
}

type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string    `json:"purePhoneNumber"` //没有区号的手机号
	CountryCode     string    `json:"countryCode"`     // 区号
	Watermark       Watermark `json:"watermark"`       // 数据水印
}

type Watermark struct {
	Timestamp int64  `json:"timestamp"` // 获取手机号的时间戳
	Appid     string `json:"appid"`     // 小程序appid
}
