package v2

import (
	"fmt"
	"strconv"
)

// 3.3余额查询接口
func (c *Client) GetAccount() (result *GetAccountResult, err error) {
	timestamp := Timestamp()
	sign := Md5Sign(c.Cfg.AppKey + timestamp + c.Cfg.AppSecret)
	err = c.Request("/Member/GetAccount", fmt.Sprintf("AppKey=%s&TimesTamp=%s&Sign=%s", c.Cfg.AppKey, timestamp, sign), &result)
	if err != nil {
		return
	}
	if result.Code != ERR_SUCCESS.Code {
		err = Err(result.Code, result.Msg)
		return
	}
	if result.Sign != Md5Sign(c.Cfg.AppKey+strconv.FormatInt(result.TimesTamp, 10)+c.Cfg.AppSecret) {
		err = ERR_SIGN
		return
	}
	return
}

type GetAccountResult struct {
	Code      int64
	Msg       string
	TimesTamp int64
	Sign      string
	Data      []*GetAccountResultData
}

type GetAccountResultData struct {
	MemberAccountCode string
	Balance           int64
}
