package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

/**
 * 5.71.1
 * 本接口提供了会员中心积点赠送的功能
 * 注:暂不支持手机号,inMode参数使用0-主动领取入账，其中1-直接入账不成功
 */
func MemberPointAcquire(c *Config, p *MemberPointAcquireParams) (r *MemberPointAcquireResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("sysId", p.SysId)
	bm.Set("openId", p.OpenId)
	bm.Set("mobile", p.Mobile)
	bm.Set("getSource", p.GetSource)
	bm.Set("transSeqId", p.TransSeqId)
	bm.Set("transTs", p.TransTs)
	bm.Set("pointAt", p.PointAt)
	bm.Set("transDigest", p.TransDigest)
	bm.Set("descCode", p.DescCode)
	bm.Set("inMode", p.InMode)
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
	resp, err := Post(c, "/memberPointAcquire", bm)
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
type MemberPointAcquireParams struct {
	SysId       string `json:"sysId"`
	OpenId      string `json:"openId"`
	Mobile      string `json:"mobile"`
	GetSource   string `json:"getSource"`
	TransSeqId  string `json:"transSeqId"`
	TransTs     string `json:"transTs"`
	PointAt     string `json:"pointAt"`
	TransDigest string `json:"transDigest"`
	DescCode    string `json:"descCode"`
	InMode      string `json:"inMode"`
}

//响应结构体
type MemberPointAcquireResult struct {
	Resp string `json:"resp"`
	Msg  string `json:"msg"`
}
