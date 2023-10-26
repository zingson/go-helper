package wxapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html

// Token 获取接口调用凭据，token有效期为7200s，开发者需要进行妥善保存。
func Token(config *Config) (rs TokenResult, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", config.Appid, config.Secret))
	if err != nil {
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return
	}
	if rs.Errcode != 0 {
		err = errors.New(strconv.FormatInt(rs.Errcode, 10) + ":" + rs.Errmsg)
		return
	}
	return
}

type TokenResult struct {
	Errcode     int64  `json:"errcode"` // 0=成功，其它失败
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // 默认7200秒
}
