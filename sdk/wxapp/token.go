package wxapp

import (
	"fmt"
)

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html

// Token 获取接口调用凭据，token有效期为7200s，开发者需要进行妥善保存。
func Token(config *Config) (rs TokenResult, err error) {
	return get[TokenResult](fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", config.Appid, config.Secret))
}

type TokenResult struct {
	Errcode     int64  `json:"errcode"` // 0=成功，其它失败
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // 默认7200秒
}
