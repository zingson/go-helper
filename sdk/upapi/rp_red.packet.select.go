package upapi

import (
	"encoding/json"
	"time"
)

/*
机构账户（红包）余额查询
查询银联红包接入方账户余额查询。
*/

// RedPacketSelect 机构账户（红包）余额查询
func RedPacketSelect(c *Config, insAccId string, backendToken func(config *Config) string) (rs *RedPacketResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Appid)
	bm.Set("insAcctId", insAccId)
	bm.Set("transNumber", Rand32())
	bm.Set("transDt", time.Now().Format("20060102"))
	bm.Set("transTm", time.Now().Format("20060102150405"))
	bm.Set("backendToken", backendToken(c))

	resp, err := Post(c, "/red.packet.select", bm)
	if err != nil {
		return
	}
	if resp.Resp != E00.Code {
		err = ErrNew(resp.Resp, resp.Msg)
		return
	}
	data, _ := json.Marshal(resp.Params)
	err = json.Unmarshal(data, &rs)
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

	rs.AcctBalance, err = Decode3DES(c.SymmetricKey, rs.AcctBalance)
	if err != nil {
		return
	}
	rs.AcctSt, err = Decode3DES(c.SymmetricKey, rs.AcctSt)
	if err != nil {
		return
	}

	return
}

type RedPacketResult struct {
	RespCode       string `json:"respCode"`
	RespMsg        string `json:"respMsg"`
	AcctSt         string `json:"acctSt"`
	AcctBalance    string `json:"acctBalance"`
	ValidBeginDtTm string `json:"validBeginDtTm"`
	ValidEndDtTm   string `json:"validEndDtTm"`
}

func (o *RedPacketResult) Json() string {
	b, _ := json.Marshal(o)
	return string(b)
}
