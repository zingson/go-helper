package upapi

import (
	"encoding/json"
)

/**
 * 5.76.1
 * 本接口提供了查询用户62VIP会员状态服务
 */
func MemberVipBalance(c *Config, p *MemberVipBalanceParams, backendToken func(config *Config) string) (r *MemberVipBalanceResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("sysId", p.SysId)
	bm.Set("openId", p.OpenId)
	bm.Set("mobile", p.Mobile)
	bm.Set("backendToken", backendToken(c))

	signature, err := UpRsaSign(bm, c.MchPrivateKey, false)
	if err != nil {
		return
	}
	bm.Set("signature", signature)
	if len(p.Mobile) > 0 {
		// 注意：签名之后再加密敏感字段
		var mobile string
		mobile, err = Encode3DES(c.SymmetricKey, p.Mobile)
		if err != nil {
			return
		}
		bm.Set("mobile", mobile)
	}
	resp, err := Post(c, "/memberVipBalance", bm)
	if err != nil {
		return
	}
	if resp.Resp != E00.Code {
		err = ErrNew(resp.Resp, resp.Msg)
		return
	}
	// 解析响应报文
	pBytes, err := json.Marshal(resp.Params)
	if err != nil {
		return
	}
	err = json.Unmarshal(pBytes, &r)
	if err != nil {
		return
	}
	return
}

// 请求参数
type MemberVipBalanceParams struct {
	SysId      string `json:"sysId"`
	OpenId     string `json:"openId"`
	Mobile     string `json:"mobile"`
	MemberType string `json:"memberType"`
}

// 响应结构体
type MemberVipBalanceResult struct {
	Resp   string                     `json:"resp"`
	Msg    string                     `json:"msg"`
	Params MemberVipBalanceResultData `json:"params"`
}

// 响应结构体参数信息
type MemberVipBalanceResultData struct {
	Status        string `json:"status"`
	MemberType    string `json:"memberType"`
	NewMember     string `json:"newMember"`
	BeginTime     string `json:"beginTime"`
	EndTime       string `json:"endTime"`
	YearValid     string `json:"yearValid"`
	SeasonValid   string `json:"seasonValid"`
	MonthValid    string `json:"monthValid"`
	HalfYearValid string `json:"halfYearValid"`
}
