package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

// MemberPointDeduct 本接口提供核销用户积点的功能
func MemberPointDeduct(c *Config, p *MemberPointDeductParams) (r *MemberPointDeductResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("sysId", p.SysId)
	bm.Set("openId", p.OpenId)
	bm.Set("mobile", p.Mobile)
	bm.Set("costSource", p.CostSource)
	bm.Set("transSeqId", p.TransSeqId)
	bm.Set("transTs", p.TransTs)
	bm.Set("pointAt", p.PointAt)
	bm.Set("transDigest", p.TransDigest)
	bm.Set("descCode", p.DescCode)
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
	if len(p.PointAt) > 0 {
		// 注意：签名之后再加密敏感字段
		var pointAt string
		pointAt, err = Encode3DES(c.SymmetricKey, p.PointAt)
		if err != nil {
			return
		}
		bm.Set("pointAt", pointAt)
	}
	resp, err := Post(c, "/memberPointDeduct", bm)
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
type MemberPointDeductParams struct {
	SysId       string `json:"sysId"`       //
	OpenId      string `json:"openId"`      //
	Mobile      string `json:"mobile"`      //
	CostSource  string `json:"costSource"`  //必填，积点消耗渠道标识
	TransSeqId  string `json:"transSeqId"`  //
	TransTs     string `json:"transTs"`     //
	PointAt     string `json:"pointAt"`     //
	TransDigest string `json:"transDigest"` //必填，积点消耗描述
	DescCode    string `json:"descCode"`    // 必填，文案代码	会员中心系统为接入方分配
}

// 响应结构体
type MemberPointDeductResult struct {
	Resp string `json:"resp"`
	Msg  string `json:"msg"`
}
