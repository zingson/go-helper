package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

/*
接入方给云闪付用户赠送专享红包。
文档： https://opentools.95516.com/applet/#/docs/develop/api-backend?id=_02040801
*/

//PointAcquire 赠送专享红包
func PointAcquire(c *Config, p *PointAcquireParams) (err error) {
	if p.AccEntityTp == "" {
		p.AccEntityTp = AETP_03
	}

	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("transSeqId", p.TransSeqId)
	bm.Set("transTs", p.TransTs)
	bm.Set("nonceStr", Rand32()[:16])
	bm.Set("insAcctId", p.InsAcctId)
	bm.Set("mobile", p.Mobile)
	bm.Set("cardNo", p.CardNo)
	bm.Set("openId", p.OpenId)
	bm.Set("acctEntityTp", string(p.AccEntityTp))
	bm.Set("pointId", p.PointId)
	bm.Set("pointAt", p.PointAt)
	bm.Set("validBeginTs", p.ValidBeginTs)
	bm.Set("validEndTs", p.ValidEndTs)
	bm.Set("busiInfo", p.BusiInfo.Json())
	bm.Set("transDigest", p.TransDigest)
	bm.Set("remark", p.Remark)
	bm.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	signature, err := UpRsaSign(bm, c.MchPrivateKey, false)
	if err != nil {
		return
	}
	bm.Set("signature", signature)

	if p.Mobile != "" {
		// 注意：签名之后再加密敏感字段
		var mobile string
		mobile, err = Encode3DES(c.SymmetricKey, p.Mobile)
		if err != nil {
			return
		}
		bm.Set("mobile", mobile)
	}
	if p.CardNo != "" {
		// 注意：签名之后再加密敏感字段
		var cardNo string
		cardNo, err = Encode3DES(c.SymmetricKey, p.CardNo)
		if err != nil {
			return
		}
		bm.Set("cardNo", cardNo)
	}

	resp, err := Post(c, "/point.acquire", bm)
	if err != nil {
		return
	}

	if resp.Resp != E00.Code {
		e, ok := gpup[resp.Resp]
		if ok {
			err = e
			return
		}
		err = ErrNew(resp.Resp, resp.Msg)
		return
	}
	return
}

type PointAcquireParams struct {
	TransSeqId   string      // 必填 交易请求流水
	TransTs      string      // 必填 交易时间，格式: yyyymmddhhmmss
	InsAcctId    string      // 必填 机构账户代码，最大32位，对应云闪付小程序开放平台配置：营销能力包-红包接入方账户
	PointId      string      // 必填 积分 id ，最大 64 位,对应云闪付小程序开放平台配置：营销能力包-专享红包活动编码
	PointAt      string      // 必填 积分额，最大 12 位,单位以分计算
	BusiInfo     *BusiInfo   // 必填 业务信息，最大 1024 位，包含自定义活动 ID ,活动名称两个小字段，格式为 {" campaignId ":"活动 ID "," campaignName ":"活动名称"}，ps ：活动ID确保唯一
	TransDigest  string      // 必填 交易摘要，最大200位,赠送积分说明，用于前台展示
	AccEntityTp  AccEntityTp // 必填 账户主体类型， 2 位，可选： 01 -手机号 02 -卡号 03 -用户（三选一） 。默认03。
	Mobile       string      // 选填 交易手机号，若上送，（使用 symmetricKey 对称加密，内容为 base64格式 ）
	CardNo       string      // 选填 交易卡号，若上送，（使用 symmetricKey 对称加密，内容为 base64格式 ）
	OpenId       string      // 选填 用户唯一标识
	ValidBeginTs string      // 选填 有效起始时间，格式：yyyymmddhhmmss
	ValidEndTs   string      // 选填 有效截止时间，格式：yyyymmddhhmmss
	Remark       string      // 选填 备注，最大 512 位
}

type BusiInfo struct {
	CampaignId   string `json:"campaignId"`   // 必填 活动ID
	CampaignName string `json:"campaignName"` // 必填 活动名称
}

func (o *BusiInfo) Json() string {
	b, _ := json.Marshal(o)
	return string(b)
}
