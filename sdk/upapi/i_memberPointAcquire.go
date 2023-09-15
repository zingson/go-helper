package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

// MemberPointAcquire 本接口提供了会员中心积点赠送的功能
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

// 请求参数
type MemberPointAcquireParams struct {
	SysId       string `json:"sysId"`       //必填，系统ID 会员中心系统为接入方分配
	OpenId      string `json:"openId"`      //openid与手机号二选一
	Mobile      string `json:"mobile"`      //
	GetSource   string `json:"getSource"`   //必填。 积点获取渠道标识，会员中心系统为接入方分配
	TransSeqId  string `json:"transSeqId"`  //必填，流水号
	TransTs     string `json:"transTs"`     //必填，交易时间 t14
	PointAt     string `json:"pointAt"`     //必填，积点值
	TransDigest string `json:"transDigest"` //必填，积点赠送描述
	DescCode    string `json:"descCode"`    //必填，文案代码，会员中心系统为接入方分配
	InMode      string `json:"inMode"`      //必填，是否直接入账 0-主动领取入账  1-直接入账
}

// 响应结构体
type MemberPointAcquireResult struct {
	Resp string `json:"resp"`
	Msg  string `json:"msg"`
}
