package upapi

import (
	"encoding/json"
	"time"
)

// PointBalance 专享红包余额查询
// 根据专享红包活动 id 查询用户的该专享红包余额。
func PointBalance(c *Config, p *PointBalanceParams, backendToken func(config *Config) string) (rs *PointBalanceResult, err error) {
	if p.AccessId == "" {
		p.AccessId = "UP"
	}

	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("pointId", p.PointId)
	bm.Set("acctEntityTp", string(p.AcctEntityTp))
	bm.Set("transTs", time.Now().Format("20060102150405"))
	bm.Set("accessId", p.AccessId)
	bm.Set("transSeq", Rand32())
	bm.Set("remark", "")
	bm.Set("backendToken", backendToken(c))

	accEntityTp, err := Encode3DES(c.SymmetricKey, p.AcctEntityId)
	if err != nil {
		return
	}
	bm.Set("acctEntityId", accEntityTp)

	resp, err := Post(c, "/point.balance", bm)
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

	data, _ := json.Marshal(resp.Params)
	err = json.Unmarshal(data, &rs)
	if err != nil {
		return
	}

	return
}

type PointBalanceParams struct {
	PointId      string      //必填 专享活动活动id，对应云闪付小程序开放平台配置：营销能力包-专享红包活动编码
	AcctEntityTp AccEntityTp //必填 账户主体类型：01 -手机号；02 -卡号；03 -用户，即 openId
	AcctEntityId string      //必填 账户主体值，类型填写 01 ，则对应填写具体手机号；类型填写 02 ，则对应填写具体卡号；类型填写 03 ，则对应填写具体 openId （ 3 种类型值均需要使用 symmetricKey 对称密钥加密，内容为 base64 格式）
	AccessId     string      //必填 专享红包来源： UP -银联； SC -其他
}

type PointBalanceResult struct {
	AcctSt          string `json:"acctSt"`          //账户状态
	AvlBalance      int64  `json:"avlBalance"`      //可用金额
	FrozenBalance   string `json:"frozenBalance"`   //冻结金额
	OperBalance     string `json:"operBalance"`     //消费金额
	ExpireBalance   string `json:"expireBalance"`   //过期金额
	TodeductBalance string `json:"todeductBalance"` //待抵扣金额
	AcctOpenTp      string `json:"acctOpenTp"`      //开户方式
	AcctOpenDt      string `json:"acctOpenDt"`      //开户日期
	ReservedField   string `json:"reservedField"`   //保留字段
	RecUpdTs        string `json:"recUpdTs"`        //记录更新时间
}
