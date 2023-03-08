package bytedanceapp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session

// 登录凭证校验。通过 login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
func Jscode2session(cfg *Config, code string) (data *Jscode2sessionResultData, err error) {
	bm := map[string]string{
		"appid":  cfg.Appid,
		"secret": cfg.Secret,
		"code":   code,
	}
	b, _ := json.Marshal(&bm)
	resp, err := http.Post("https://developer.toutiao.com/api/apps/v2/jscode2session", "application/json", bytes.NewReader(b))
	if err != nil {
		return
	}
	rbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r *Jscode2sessionResult
	err = json.Unmarshal(rbytes, &r)
	if err != nil {
		return
	}
	if r.ErrNo != 0 {
		err = errors.New(strconv.FormatInt(r.ErrNo, 10) + ":" + r.ErrTips)
		return
	}
	data = r.Data
	return
}

type Jscode2sessionResult struct {
	ErrNo   int64                     `json:"err_no"`
	ErrTips string                    `json:"err_tips"`
	Data    *Jscode2sessionResultData `json:"data"`
}

type Jscode2sessionResultData struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

/*
errcode 的合法值
-1	系统繁忙，此时请开发者稍候再试
0	请求成功
40029	code 无效
45011	频率限制，每个用户每分钟100次
*/
