package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

// CouponDownload 5.8.9  赠送优惠券 <coupon.download>
func CouponDownload(c *Config, p *CouponDownloadParams) (r *CouponDownloadResult, err error) {
	if p.AcctEntityTp == "" {
		p.AcctEntityTp = "03"
	}

	bm := NewBodyMap()
	bm.Set("appId", c.Appid)                //是 接入方的唯一标识
	bm.Set("transSeqId", p.TransSeqId)      //是 交易流水号,不重复，最大64位
	bm.Set("transTs", p.TransTs)            //是 请求日期, 格式yyyyMMdd，如：20191227
	bm.Set("couponId", p.CouponId)          //是 优惠券id
	bm.Set("mobile", p.Mobile)              //否 交易手机号，若上送，（使用symmetricKey 对称加密，内容 为base64格式）
	bm.Set("cardNo", p.CardNo)              //否 交易卡号，若上送，（使用symmetricKey 对称加密，内容为 base64格式）
	bm.Set("openId", p.OpenId)              //否 用户唯一标识
	bm.Set("acctEntityTp", p.AcctEntityTp)  //是 营销活动配置的赠送维度（参见营销平台活动配置），2位， 可选：01=手机 02-卡号 03-用户（二选一） 赠送维度为卡号时，则cardNo必填； 赠送维度为用户时，则openId，mobile, cardNo三选一上送
	bm.Set("couponNum", p.CouponNum)        //是  优惠券数量
	bm.Set("nonceStr", GetRandomString(16)) //是	生成签名的随机串
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
			return nil, err
		}
		bm.Set("mobile", mobile)
	}
	if p.CardNo != "" {
		// 注意：签名之后再加密敏感字段
		var cardNo string
		cardNo, err = Encode3DES(c.SymmetricKey, p.CardNo)
		if err != nil {
			return nil, err
		}
		bm.Set("cardNo", cardNo)
	}

	resp, err := Post(c, "/coupon.download", bm)
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

	// 解析响应报文
	pbytes, err := json.Marshal(resp.Params)
	if err != nil {
		return
	}
	err = json.Unmarshal(pbytes, &r)
	if err != nil {
		return
	}
	return
}

type CouponDownloadParams struct {
	TransSeqId   string // 交易流水
	TransTs      string // 请求日期
	CouponId     string
	CouponNum    int64
	AcctEntityTp string //赠送维度 默认03
	Mobile       string //Mobile CardNo OpenId 三选一
	CardNo       string
	OpenId       string
}

type CouponDownloadResult struct {
	TransSeqId string
	CouponId   string
}
