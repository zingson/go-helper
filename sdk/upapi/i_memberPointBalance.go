package upapi

import (
	"encoding/json"
)

// MemberPointBalance 本接口提供会员中心积点余额查询的功能
func MemberPointBalance(c *Config, p *MemberPointBalanceParams) (r *MemberPointBalanceResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("sysId", p.SysId)
	bm.Set("openId", p.OpenId)
	bm.Set("mobile", p.Mobile)
	bm.Set("transSeqId", p.TransSeqId)
	bm.Set("transTs", p.TransTs)
	bm.Set("backendToken", p.BackendToken)
	if len(p.Mobile) > 0 {
		// 注意：签名之后再加密敏感字段
		var mobile string
		mobile, err = Encode3DES(c.SymmetricKey, p.Mobile)
		if err != nil {
			return
		}
		bm.Set("mobile", mobile)
	}
	resp, err := Post(c, "/memberPointBalance", bm)
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
type MemberPointBalanceParams struct {
	SysId        string `json:"sysId"`
	OpenId       string `json:"openId"`
	Mobile       string `json:"mobile"`
	TransSeqId   string `json:"transSeqId"`
	TransTs      string `json:"transTs"`
	BackendToken string `json:"backendToken"`
}

// 响应结构体
type MemberPointBalanceResult struct {
	Resp   string                       `json:"resp"`
	Msg    string                       `json:"msg"`
	Params MemberPointBalanceResultData `json:"params"`
}

// 响应结构体参数信息
type MemberPointBalanceResultData struct {
	AccSt           string `json:"acctSt"`
	AvlBalance      string `json:"avlBalance"`
	FrozenBalance   string `json:"frozenBalance"`
	OperBalance     string `json:"operBalance"`
	ExpireBalance   string `json:"expireBalance"`
	ToDeductBalance string `json:"todeductBalance"`
	AccOpenTp       string `json:"acctOpenTp"`
	AcctOpenDt      string `json:"acctOpenDt"`
	ReservedField   string `json:"reservedField"`
	RecUpdTs        string `json:"recUpdTs"`
}
