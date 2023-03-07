package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

/**
 * 5.69.1
 * 本接口提供给用户赠送62VIP会员的功能
 */
func MemberVipAcquire(c *Config, p *MemberVipAcquireParams) (r *MemberVipAcquireResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("sysId", p.SysId)
	bm.Set("openId", p.OpenId)
	bm.Set("mobile", p.Mobile)
	bm.Set("memberType", p.MemberType)
	bm.Set("transSeqId", p.TransSeqId)
	bm.Set("transTs", p.TransTs)
	bm.Set("isLimit", p.IsLimit)
	bm.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	bm.Set("nonceStr", GetRandomString(10))
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
	resp, err := Post(c, "/memberVipAcquire", bm)
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

//请求参数
type MemberVipAcquireParams struct {
	SysId      string `json:"sysId"`
	OpenId     string `json:"openId"`
	Mobile     string `json:"mobile"`
	MemberType string `json:"memberType"`
	TransSeqId string `json:"transSeqId"`
	TransTs    string `json:"transTs"`
	IsLimit    string `json:"isLimit"`
}

//响应结构体
type MemberVipAcquireResult struct {
	Resp string `json:"resp"`
	Msg  string `json:"msg"`
}
